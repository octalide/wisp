# wisp

Wisp (short for whisper) is a very tiny, multithread safe event subsystem. Under 100 lines of code!!

Wisp operates on an extremely simple event/handler platform and uses "tags" (string identifiers) for event filtering.
Handlers have the option to consume events by switching to blocking mode and returning a value of `true` from the callback.
