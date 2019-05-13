-module(my_chat_client).
-export([
    start/0
]).

-define(INFO(A1, A2), io:format(lists:concat(["mod:~w,line:~w,\t", A1]), [?MODULE, ?LINE | A2])).

start() ->
%%    case gen_tcp:connect("localhost", 6789, [binary, {packet, 0}]) of
%%    case gen_tcp:connect("192.168.1.69", 6789, [binary, {packet, 0}]) of
    case gen_tcp:connect("192.168.1.69", 6789, [binary, {packet, 4}]) of
        {ok, Socket} ->
            ?INFO("connect succeed!:~w~n", [Socket]),
            spawn(fun() -> loop(Socket) end),
            send_sth(Socket);
        Other ->
            ?INFO("other:~w~n", [Other])
    end.

loop(Socket) ->
    receive
        Any ->
            ?INFO("any:~w~n", [Any]),
            loop(Socket)
    end.

send_sth(Socket) ->
    Str = "hello",
    gen_tcp:send(Socket, term_to_binary(Str)).
