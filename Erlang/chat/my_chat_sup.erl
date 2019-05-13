-module(my_chat_sup).
-behavior(supervisor).
-export([
    start_link/0,
    init/1
]).

start_link() ->
    io:format("my_chat_sup start...~n"),
    supervisor:start_link({local, ?MODULE}, ?MODULE, []).

init([]) ->
    io:format("my_chat_sup init...~n"),
    {ok, {
        {one_for_one, 3, 10},
        []
    }}.
