-module(v7_ews).
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
	    loop(Sock);
	    
	Any ->
	    io:format("receive any: ~p~n", [Any])
    end.

loop(Sock) ->
    receive
	{tcp, Sock, Data} ->
	    FrameOpcode = 0,
	    try
	        Text = process_buffer(Data, FrameOpcode, Sock, <<>>),
		io:format("parse client data text: ~p~n", [Text])
	    catch
		throw: X ->
		    io:format("catch error: ~p~n", [X])
	    end,

	    loop(Sock);
	{tcp_closed, Sock} ->
	    ok;
	_Any ->
	    loop(Sock)
    end.


process_buffer(Data, FrameOpcode, Sock, Payload0) when is_list(Data) ->
    process_buffer(list_to_binary(Data), FrameOpcode, Sock, Payload0);

process_buffer(<< Fin:1, _RSV:3, Opcode:4, _Masked:1, PayloadLen:7, Rest/bits >> = Bin, FrameOpcode, Sock, Payload0) ->
    FrameOpcode2 = 
    if 
	Fin =:= 0 ->
	    FrameOpcode1 = Opcode bor FrameOpcode,
	    FrameOpcode1;
	true ->
	    FrameOpcode
    end,

    ByteSizeBin = byte_size(Bin),
    (PayloadLen > 125) andalso (ByteSizeBin < 8) andalso throw({error, package_size_small}),

    {RealPayloadLen, MaskKey, Rest2} = 
    if 
	PayloadLen =:= 126 ->
	    << PayloadLen1:16, MaskKey0:32, Rest1/bits >> = Rest,
	    {PayloadLen1, MaskKey0, Rest1};
	PayloadLen =:= 127 ->
	    << PayloadLen1:64, MaskKey0:32, Rest1/bits >> = Rest,
	    {PayloadLen1, MaskKey0, Rest1};
	true ->
	    << MaskKey0:32, Rest1/bits >> = Rest,
	    {PayloadLen, MaskKey0, Rest1}
    end,
    ByteSizeRest2 = byte_size(Rest2),
    %% TODO 拼接下一个tcp分片
    (ByteSizeRest2 < RealPayloadLen) andalso throw({error, package_rest_not_enough}),

    << MaskPayload:RealPayloadLen/binary, RealRest/bits >> = Rest2,
    Payload = data_unmask(MaskPayload, MaskKey, <<>>),
    Payload1 = <<Payload0/binary, Payload/binary>>,

    if
	Fin =:= 0 ->
	    %% TODO 拼接下一个websocket包
    	    process_buffer(RealRest, FrameOpcode2, Sock, Payload1);
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
