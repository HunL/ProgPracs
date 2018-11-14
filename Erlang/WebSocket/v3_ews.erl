-module(v3_ews).
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

	    Text = parse_client_data_check(Data, Sock),
	    io:format("parse client data text: ~p~n", [Text]),
	    send_data(Sock, ["client data: ", Text]),

	    loop(Sock);

	{tcp_closed, Sock} ->
	    io:format("tcp closed.~n", []),
	    ok;

	Any ->
	    io:format("loop Any: ~p~n", [Any]),
	    loop(Sock)
    end.


parse_client_data_check(Data, Sock) when is_list(Data) ->
    parse_client_data_check(list_to_binary(Data), Sock);


parse_client_data_check(<< Fin:1, RSV:3, Opcode:4, Masked:1, PayloadLen:7, Rest/bits >>, Sock) ->
    io:format("client data fin: ~p~n", [Fin]),
    io:format("client data rsv: ~p~n", [RSV]),
    io:format("client data opcode: ~p~n", [Opcode]),
    io:format("client data masked: ~p~n", [Masked]),
    io:format("client data PayloadLen: ~p~n", [PayloadLen]),

    if
	%% 断开
%%	(Fin =:= 0) andalso (RSV =:= 0) ->
%%	    io:format("close socket!~n", []),
%%	    gen_tcp:close(Sock);
	%% 断开
	Opcode =:= 16#8 ->
	    io:format("close socket!~n", []),
	    gen_tcp:close(Sock);
	%% 没有掩码，断开
	Masked =:= 16#0 ->
	    io:format("close socket!~n", []),
	    gen_tcp:close(Sock);
	%% 不支持分片
	Opcode =:= 16#0 ->
	    io:format("close socket!~n", []),
	    gen_tcp:close(Sock);
	%% 二进制帧
	Opcode =:= 16#2 ->
	    io:format("received binary opcode: ~p~n", [Opcode]),
	    <<>>;
	%% text
	Opcode =:= 16#1 ->
	    io:format("received text opcode: ~p~n", [Opcode]),
	    parse_client_data(PayloadLen, Rest, Sock);
	    %% << Value:RealPayloadLen/binary, _/bits >> = Rest1,
	    %% Text = data_unmask(Value, MaskKey, <<>>),
	    %% Text;
	    %% <<>>;
	%% ping
	Opcode =:= 16#9 ->
	    io:format("received ping opcode: ~p~n", [Opcode]),
	    <<"pong">>;
	%% 不支持pong
	Opcode =:= 16#A ->
	    %% io:format("close socket!~n", []),
	    %% gen_tcp:close(Sock);
	    io:format("unsupported pong opcode: ~p~n", [Opcode]),
	    <<>>;
	true ->
	    io:format("unsupported opcode: ~p~n", [Opcode]),
	    <<>>
    end.

parse_client_data(PayloadLen, Rest, _Sock) ->
    if 
	PayloadLen =:= 126 ->
	    << RealPayloadLen:16, MaskKey:32, Rest1/bits >> = Rest,
	    io:format("16bit realpayloadlen: ~p~n", [RealPayloadLen]),
	    Restiosize = iolist_size(Rest),
	    io:format("rest iosize: ~p~n", [Restiosize]),
	    << Value:RealPayloadLen/binary, _/bits >> = Rest1,
	    io:format("realpayloadlen * binary size: ~p~n", [RealPayloadLen * 8]),
	    io:format("bit_size value: ~p~n", [bit_size(Value)]),
	    io:format("byte_size value: ~p~n", [byte_size(Value)]),
	    %% << Value:RealPayloadLen/bits, _/bits >> = Rest1,
	    Text = data_unmask(Value, MaskKey, <<>>),
	    Text;
	PayloadLen =:= 127 ->
	    << RealPayloadLen:64, MaskKey:32, Rest1/bits >> = Rest,
	    io:format("64bit realpayloadlen: ~p~n", [RealPayloadLen]),
	    << Value:RealPayloadLen/binary, _/bits >> = Rest1,
	    %% << Value:RealPayloadLen/bits, _/bits >> = Rest1,
	    Text = data_unmask(Value, MaskKey, <<>>),
	    Text;
	PayloadLen < 126 ->
	    << MaskKey:32, Value:PayloadLen/binary, _/bits >> = Rest,
	    io:format("bit_size value: ~p~n", [bit_size(Value)]),
	    io:format("byte_size value: ~p~n", [byte_size(Value)]),
	    Text = data_unmask(Value, MaskKey, <<>>),
	    Text
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
