-module(ews).
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
    %% DataStr = "Conn: Upgrade\r\nUpgrade:websocket\r\nSec-Web-Ver:13\r\nSec-WebSocket-Key:abcd",
    %% test_str_tokens(DataStr, "\r\n"),
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
	    io:format("loop data: ~p~n", [Data]),
	    %% send_data(Sock, Data),
	    %% IsStr = is_list(Data),
	    %% send_data(Sock, ["receive data..."]),

	    Text = client_data(Data),
	    io:format("client_data text: ~p~n", [Text]),
	    send_data(Sock, ["client data: ", Text]),

	    loop(Sock);

	{tcp_closed, Sock} ->
	    io:format("tcp closed.~n", []),
	    ok;

	Any ->
	    io:format("loop Any: ~p~n", [Any]),
	    loop(Sock)
    end.


client_data(Data) when is_list(Data) ->
    client_data(list_to_binary(Data));
client_data(<< 1:1, 0:3, 1:4, 1:1, Len:7, MaskKey:32, Rest/bits >>) when Len < 126 ->
    << Value:Len/binary, _/bits >> = Rest,
    Text = data_unmask(Value, MaskKey, <<>>),
    Text;
client_data(_) ->
    <<>>.

%% data_unmask(Value, MaskKey, Bin) ->
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
