-module(hot_test).
-export([start/0]).
-export([do_loop/0]).

start() ->
    Pid = spawn(fun() -> do_loop() end),
    register(?MODULE, Pid).

do_loop() ->
    receive
	code_switch ->
	    io:format("receive code_switch msg~n", []),
	    ?MODULE:do_loop();
	Msg ->
	    io:format("receive msg: ~p~n", [Msg]),
    	    do_loop()
    end.

