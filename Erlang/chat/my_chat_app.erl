-module(my_chat_app).
-behavior(application).
-export([
    start/0,
    start/2,
    stop/1
]).

start() ->
    application:start(chat).

start(_, []) ->
    io:format("my_chat start...~n"),
    {ok, Pid} = my_chat_sup:start_link(),
    my_chat_server:start(),
    io:format("my_chat Main Pid:~w~n", [Pid]),
    {ok, Pid}.

stop(_) ->
    io:format("app stop!").
