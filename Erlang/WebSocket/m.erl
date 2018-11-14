-module(m).
-export([wait_msg/2]).

wait_msg(Sock, MsgData) ->
    io:format("mod m sock: ~p~n", [Sock]),
    io:format("mod m receive msg: ~p~n", [MsgData]),
    case MsgData of
        {tcp, Sock, Data} ->
	    io:format("mod m receive tcp data...~n", []),
            Key = v11_ews:get_ws_key(Data, "\r\n"),
            Challenge = base64:encode(crypto:hash(sha, << Key/binary, "258EAFA5-E914-47DA-95CA-C5AB0DC85B11" >>)),
            Handshake =
                ["HTTP/1.1 101 Switching Protocols\r\n",
                 "connection: Upgrade\r\n",
                 "upgrade: websocket\r\n",
                 "sec-websocket-accept: ", Challenge, "\r\n",
                 "\r\n",<<>>],
            gen_tcp:send(Sock, Handshake),
     
            v11_ews:send_data(Sock, ["hello man!"]),
            v11_ews:loop(Sock, <<>>);
     
         Any ->
             io:format("mod m receive any: ~p~n", [Any])
    end.
