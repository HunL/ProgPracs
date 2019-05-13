-module(my_chat_server).
-behavior(gen_server).
-export([init/1, handle_call/3, handle_cast/2, handle_info/2, terminate/2, code_change/3]).
-export([
    start/0,
    start_link/0
]).

-record(state, {}).

-define(DO_HANDLE(Request),
    try
        handle(Request)
    catch _:Reason ->
        io:format("INFO:~w~n~n, Reason: ~w~n, strace:~p", [Request, Reason, erlang:get_stacktrace()]),
        Reason
    end).

-define(INFO(A1, A2), io:format(lists:concat(["mod:~w,line:~w,\t", A1]), [?MODULE, ?LINE | A2])).

start() ->
    supervisor:start_child(my_chat_sup, {?MODULE, {?MODULE, start_link, []}, permanent, infinity, worker, [?MODULE]}).

start_link() ->
    gen_server:start_link({local, ?MODULE}, ?MODULE, [], []).

handle_call(Request, _From, State) ->
    Reply = ?DO_HANDLE(Request),
    {reply, Reply, State}.
handle_cast(Request, State) ->
    ?DO_HANDLE(Request),
    {noreply, State}.
handle_info(Request, State) ->
    ?DO_HANDLE(Request),
    {noreply, State}.

terminate(Reason, _State) ->
    ?INFO("terminate:~w~n", [Reason]),
%%    dump_db(),
%%    print_server_info
    ok.

code_change(_OldVsn, State, _Extra) ->
    {ok, State}.

handle(Func) when is_function(Func) ->
    Func();
handle({func, Func}) ->
    Func();
handle(UnkownInfo) ->
    ?INFO("MOD:~w未知消息：~w", [?MODULE, UnkownInfo]).

init([]) ->
    case gen_tcp:listen(6789, [binary, {packet, 4}, {reuseaddr, true}, {active, true}]) of
        {ok, Listen} ->
            ?INFO("listen:~w~n", [Listen]),
            spawn(fun() -> parse_connect(Listen) end);
        {error, Why} ->
            ?INFO("server start error. ~w~n", [Why])
    end,
    ?INFO("server init finished!", []),
    {ok, #state{}}.

parse_connect(Listen) ->
    case gen_tcp:accept(Listen) of
        {ok, Socket} ->
            ?INFO("socket:~w~n", [Socket]),
            spawn(fun() -> parse_connect(Listen) end),
            loop(Socket);
        {error, Why} ->
            ?INFO("accept error:~w~n", [Why])
    end.

loop(Socket) ->
    receive
        {tcp, Socket, Bin} ->
            Term = binary_to_term(Bin),
            ?INFO("term:~w~n", [Term]),
            ok;
        Other ->
            ?INFO("other:~w~n", [Other]),
            ok
    end.
