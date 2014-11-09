#Gron - a cron in go.

## Goals
The goal of Gron is to create a simple small go library able to recreate some of the functionality of cron ( http://en.wikipedia.org/wiki/Cron ).

Gron will include a number of different modules.

1. Parsing - an easy to use set of methods for parsing strings into recurring schedules: "every friday do X", "the last day of every month do Y".

2. Scheduling - based on a schedule struct, understand if a command should be "run", and the next epoch timestamp it should be run.

3. Running - a goroutine to start that will handle parsing, scheduling and invoking of schedule structs (stored in memory).

