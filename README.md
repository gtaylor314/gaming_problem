# Gaming Problem
Given two strings in format HH:MM, find the number of playable 15 minute games

## Requirements
Games only run on the hour/quarter hour and require the full 15 minutes to play

For example: 12:10 - 13:15 has four playable games 

For example: 12:01 - 12:16 has no playable games since the game starts at 12:00

Time is in the 24-hour format

If time A = time B then we treat this as being within the same minute, meaning no playable games
If time A > time B then game play has rolled over into the next day