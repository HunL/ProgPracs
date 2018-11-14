-module(v8_ews).
-export([start/0]).

start() ->
    {ok, ListenSock} = gen_tcp:listen(8000, [{packet, 0}, {active, true}, {reuseaddr, true}]),
    accept_handle(ListenSock).

accept_handle(ListenSock) ->
    case gen_tcp:accept(ListenSock) of
	{ok, Sock} ->
	    spawn(fun() -> accept_handle(ListenSock) end),
	    wait_msg(Sock);
	{error, _} ->
	    start()
    end.

wait_msg(Sock) ->
    receive 
	{tcp, Sock, Data} ->
	    Key = get_ws_key(Data, "\r\n"),
 	    Challenge = base64:encode(crypto:hash(sha, << Key/binary, "258EAFA5-E914-47DA-95CA-C5AB0DC85B11" >>)),
            Handshake =
 		    ["HTTP/1.1 101 Switching Protocols\r\n",
 		     "connection: Upgrade\r\n",
 		     "upgrade: websocket\r\n",
 		     "sec-websocket-accept: ", Challenge, "\r\n",
 		     "\r\n",<<>>],
            gen_tcp:send(Sock, Handshake),

	    send_data(Sock, ["hello man!"]),
	    loop(Sock, <<>>);
	    
	Any ->
	    io:format("receive any: ~p~n", [Any])
    end.

loop(Sock, Data0) ->
    io:format("loop data0: ~p~n", [Data0]),
    receive
	{tcp, Sock, Data} ->
	    FrameOpcode = 0,
	    try
		if 
		    Data0 == <<>> ->
			Text = process_buffer(Data, FrameOpcode, Sock, <<>>),
			io:format("111 parse client data text: ~p~n", [Text]),
	                loop(Sock, <<>>);
		    true ->
			io:format("data0 is list: ~p~n", [is_list(Data0)]),
			io:format("data0 is binary: ~p~n", [is_binary(Data0)]),
			io:format("data is list: ~p~n", [is_list(Data)]),
			io:format("data is binary: ~p~n", [is_binary(Data)]),

			BinData = list_to_binary(Data),
			NewData = <<Data0/binary, BinData/binary>>,
			Text = process_buffer(NewData, FrameOpcode, Sock, <<>>),
			io:format("nnn parse client data text: ~p~n", [Text]),
%	                loop(Sock, Data0)
	                loop(Sock, <<>>)
		end
	    catch
		throw: X ->
		    io:format("catch error: ~p~n", [X])
	    end;

	{tcp_closed, Sock} ->
	    ok;
	_Any ->
	    loop(Sock, Data0)
    end.


process_buffer(Data, FrameOpcode, Sock, Payload0) when is_list(Data) ->
    process_buffer(list_to_binary(Data), FrameOpcode, Sock, Payload0);

process_buffer(<< Fin:1, RSV:3, Opcode:4, Masked:1, PayloadLen:7, Rest/bits >> = Bin, FrameOpcode, Sock, Payload0) ->
    io:format("~n~n~nprocess buffer client data fin: ~p~n", [Fin]),
    io:format("process buffer client data rsv: ~p~n", [RSV]),
    io:format("process buffer client data opcode: ~p~n", [Opcode]),
    io:format("process buffer client data masked: ~p~n", [Masked]),
    io:format("process buffer client data PayloadLen: ~p~n", [PayloadLen]),

    io:format("process buffer client data FrameOpcode: ~p~n", [FrameOpcode]),

    FrameOpcode2 = 
    if 
	Fin =:= 0 ->
	    FrameOpcode1 = Opcode bor FrameOpcode,
	    FrameOpcode1;
	true ->
	    FrameOpcode
    end,

    ByteSizeBin = byte_size(Bin),
    io:format("bytesize of bin: ~p~n", [ByteSizeBin]),

    (PayloadLen > 125) andalso (ByteSizeBin < 8) andalso throw({error, package_size_small}),

    {RealPayloadLen, MaskKey, Rest2} = 
    if 
	PayloadLen =:= 126 ->
	    << PayloadLen1:16, MaskKey0:32, Rest1/bits >> = Rest,
	    io:format("16bit realpayloadlen:--------------> ~p~n", [PayloadLen1]),
	    {PayloadLen1, MaskKey0, Rest1};
	PayloadLen =:= 127 ->
	    << PayloadLen1:64, MaskKey0:32, Rest1/bits >> = Rest,
	    io:format("64bit realpayloadlen:--------------> ~p~n", [PayloadLen1]),
	    {PayloadLen1, MaskKey0, Rest1};
	true ->
	    io:format("other payload len:--------------> ~p~n", [PayloadLen]),
	    << MaskKey0:32, Rest1/bits >> = Rest,
	    {PayloadLen, MaskKey0, Rest1}
    end,
    io:format("maskkey: ~p~n~n", [MaskKey]),
    ByteSizeRest2 = byte_size(Rest2),
    io:format("byte size of rest2:-----------------> ~p~n", [ByteSizeRest2]),

    %% TODO 拼接下一个tcp分片
%    (ByteSizeRest2 < RealPayloadLen) andalso throw({error, package_rest_not_enough}),
    if
	(ByteSizeRest2 < RealPayloadLen) ->
	    io:format("package_rest_not_enough, continue next tcp data~n", []),
%    	    process_buffer(Rest2, FrameOpcode2, Sock, Payload0);
	    loop(Sock, Bin);
	true ->
	    ok
    end,

    << MaskPayload:RealPayloadLen/binary, RealRest/bits >> = Rest2,
    Payload = data_unmask(MaskPayload, MaskKey, <<>>),
    Payload1 = <<Payload0/binary, Payload/binary>>,

    io:format("Fin: ~p~n", [Fin]),
    if
	Fin =:= 0 ->
	    %% TODO 拼接下一个websocket包
    	    process_buffer(Bin, FrameOpcode2, Sock, Payload1);
	Fin =:= 1 ->
	    Opcode1 = Opcode bor FrameOpcode2,
	    handle_frame(Opcode1, Payload1, Sock),
	    Payload1;
	true ->
	    Payload1
    end.


handle_frame(Opcode, Payload, Sock) ->
    if
	%% 断开
	Opcode =:= 16#8 ->
	    gen_tcp:close(Sock);
	%% text
	Opcode =:= 16#1 ->
	    send_data(Sock, ["client data: ", Payload]);
	%% 二进制帧，直接交付
	Opcode =:= 16#2 ->
	    <<>>;
	%% ping，发送pong响应
	Opcode =:= 16#9 ->
	    <<"pong">>;
	%% pong，不处理
	Opcode =:= 16#A ->
	    <<>>;
	true ->
	    io:format("unsupported opcode: ~p~n", [Opcode]),
	    <<>>
    end.
    
data_unmask(<<>>, _, Unmasked) ->
    Unmasked;
data_unmask(<< O:32, Rest/bits >>, MaskKey, Acc) ->
    T = O bxor MaskKey,
    data_unmask(Rest, MaskKey, << Acc/binary, T:32 >>);
data_unmask(<< O:24 >>, MaskKey, Acc) ->
    << MaskKey2:24, _:8 >> = << MaskKey:32 >>,
    T = O bxor MaskKey2,
    << Acc/binary, T:24 >>;
data_unmask(<< O:16 >>, MaskKey, Acc) ->
    << MaskKey2:16, _:16 >> = << MaskKey:32 >>,
    T = O bxor MaskKey2,
    << Acc/binary, T:16 >>;
data_unmask(<< O:8 >>, MaskKey, Acc) ->
    << MaskKey2:8, _:24 >> = << MaskKey:32 >>,
    T = O bxor MaskKey2,
    << Acc/binary, T:8 >>.

get_ws_key(Data, SeparatorList) ->
    AAA = string:tokens(Data, SeparatorList),
    Fun = fun(S) -> lists:prefix("Sec-WebSocket-Key:", S) end,
    BBB = lists:filter(Fun, AAA),
    HHH = hd(BBB),
    CCC = string:tokens(HHH, ": "),
    LLL = lists:last(CCC),
    list_to_binary(LLL).

send_data(Sock, Payload) ->
    Len = iolist_size(Payload),
    BinLen = payload_len(Len),
    gen_tcp:send(Sock, [<<1:1, 0:3, 1:4, 0:1, BinLen/bits>>, Payload]).


payload_len(N) ->
    case N of
        N when N =< 125 -> << N:7 >>;
	N when N =< 16#ffff -> << 126:7, N:16 >>;
	N when N =< 16#7fffffffffffffff -> << 127:7, N:64 >>
    end.
