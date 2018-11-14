-module(v4_ews).
-export([start/0]).

start() ->
    io:format("111~n", []),
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
	    io:format("~n~n~nloop data: ~p~n", [Data]),

%%	    Text = parse_client_data_check(Data, Sock),
	    FrameOpcode = 0,
	    Text = process_buffer(Data, FrameOpcode, Sock),
	    io:format("parse client data text: ~p~n", [Text]),
%%	    send_data(Sock, ["client data: ", Text]),

	    loop(Sock);

	{tcp_closed, Sock} ->
	    io:format("tcp closed.~n", []),
	    ok;

	Any ->
	    io:format("loop Any: ~p~n", [Any]),
	    loop(Sock)
    end.


process_buffer(Data, FrameOpcode, Sock) when is_list(Data) ->
    process_buffer(list_to_binary(Data), FrameOpcode, Sock);

process_buffer(<< Fin:1, RSV:3, Opcode:4, Masked:1, PayloadLen:7, Rest/bits >> = Bin, FrameOpcode, Sock) ->
    io:format("process buffer client data fin: ~p~n", [Fin]),
    io:format("process buffer client data rsv: ~p~n", [RSV]),
    io:format("process buffer client data opcode: ~p~n", [Opcode]),
    io:format("process buffer client data masked: ~p~n", [Masked]),
    io:format("process buffer client data PayloadLen: ~p~n", [PayloadLen]),

    if 
	Fin =:= 0 ->
	    io:format("fin == 0", []),
	    %% TODO test
	    FrameOpcode = (Opcode == 1) orelse (FrameOpcode == 0),
	    io:format("fin frameopcode: ~p~n", [FrameOpcode]),
	    ok;
	true ->
	    io:format("fin===========~n~n", []),
	    ok
    end,

    ByteSizeBin = byte_size(Bin),
    io:format("bytesize of bin: ~p~n", [ByteSizeBin]),

    if 
	PayloadLen > 125 ->
	    if 
		byte_size(Bin) < 8 ->
		    %% return error
		    io:format("error error error~n", []),
		    ok;
		true ->
	    	    io:format("bin byte > 125 size ok~n", []),
		    ok
	    end;
	true ->
	    io:format("bin byte size ok~n", []),
	    ok
    end,

    {RealPayloadLen, Payload, RealRest} = 
    if 
	PayloadLen =:= 126 ->
	    io:format("payload len is 126~n", []),
	    << PayloadLen1:16, MaskKey:32, Rest1/bits >> = Rest,
	    io:format("16bit realpayloadlen: ~p~n", [PayloadLen1]),


	    ByteSizeRest1 = byte_size(Rest1),
	    io:format("byte size of rest1: ~p~n", [ByteSizeRest1]),
	    if 
		ByteSizeRest1 + 4 < PayloadLen1 ->
		    io:format("rest size error error error~n", []),
		    ok;
		true ->
		    io:format("rest size ok ok ok~n", []),
		    ok
	    end,


	    << Value:PayloadLen1/binary, _/bits >> = Rest1,



	    %% TODO test
%%	    if 
%%		MaskKey =:= 1 ->
%%	    	    Text = data_unmask(Value, MaskKey, <<>>),
%%		    ok;
%%		true ->
%%		    ok
%%	    end,
    	    Text = data_unmask(Value, MaskKey, <<>>),
	    {PayloadLen1, Text, Rest1};
	PayloadLen =:= 127 ->
	    io:format("payload len is 127~n", []),
	    << PayloadLen1:64, MaskKey:32, Rest1/bits >> = Rest,
	    io:format("64bit realpayloadlen: ~p~n", [PayloadLen1]),
	    << Value:PayloadLen1/binary, _/bits >> = Rest1,
	    Text = data_unmask(Value, MaskKey, <<>>),
	    {PayloadLen1, Text, Rest1};
	true ->
	    io:format("other payload len: ~p~n", [PayloadLen]),
	    << MaskKey:32, Value:PayloadLen/binary, Rest1/bits >> = Rest,
            Text = data_unmask(Value, MaskKey, <<>>),
	    {PayloadLen, Text, Rest1}
    end,
    ByteSizeRealRest = byte_size(RealRest),
    io:format("byte size of real rest: ~p~n", [ByteSizeRealRest]),

    io:format("payload value: ~p~n~n", [Payload]),
	    
    if 
	ByteSizeRealRest + 4 < RealPayloadLen ->
	    io:format("rest size error error error~n", []),
	    ok;
	true ->
	    ok
    end,
	    
    ByteLength = byte_size(Payload),
    io:format("byte length of payload: ~p~n", [ByteLength]),

    if
	Fin =:= 0 ->
	    io:format("server detect fragment, sizeof payload: ~p~n", [ByteLength]);
	Fin =:= 1 ->
	    io:format("done join all data!~n", []),
	    handle_frame(Opcode, Payload, Sock),
	    ok;
	true ->
	    ok
    end,

    Payload.

handle_frame(Opcode, Payload, Sock) ->
    io:format("client data opcode: ~p~n", [Opcode]),

    if
	%% 断开
	Opcode =:= 16#8 ->
	    io:format("close socket!~n", []),
	    gen_tcp:close(Sock);
	%% text
	Opcode =:= 16#1 ->
	    io:format("received text opcode: ~p~n", [Opcode]),
	    send_data(Sock, ["client data: ", Payload]);
	%% 二进制帧，直接交付
	Opcode =:= 16#2 ->
	    io:format("received binary opcode: ~p~n", [Opcode]),
	    <<>>;
	%% ping，发送pong响应
	Opcode =:= 16#9 ->
	    io:format("received ping opcode: ~p~n", [Opcode]),
	    <<"pong">>;
	%% pong，不处理
	Opcode =:= 16#A ->
	    io:format("unsupported pong opcode: ~p~n", [Opcode]),
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
    io:format("send data, len: ~p~n", [Len]),
    BinLen = payload_len(Len),
    io:format("send data, binlen: ~p~n", [BinLen]),
    gen_tcp:send(Sock, [<<1:1, 0:3, 1:4, 0:1, BinLen/bits>>, Payload]).


payload_len(N) ->
    case N of
        N when N =< 125 -> << N:7 >>;
	N when N =< 16#ffff -> << 126:7, N:16 >>;
	N when N =< 16#7fffffffffffffff -> << 127:7, N:64 >>
    end.
