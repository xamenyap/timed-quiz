# Timed Quiz
Timed Quiz is a game of quiz written in Go running in a console terminal. The quiz contains text-based questions and answers by reading from `problems.csv` taken in 60 seconds. Both file path and quiz duration can be customized by passing certain flags when running the program. Check [here](#available-options) for all available flags

## Prerequisites
- Go v1.16

## Available Options
- `-h` displays help of all other options

Example:
```
go build . && ./timed-quiz -h
```

- `-f` to supply your own quiz file path to the program

Example:
```
go build . && ./timed-quiz -f ~/Downloads/my-downloaded-quiz.csv
```

Note that each question in your own supplied quiz file must have exactly 1 question and 1 answer on the same line (separated by a comma) like this
```
question 1, answer 1
question 2, answer 2
question 3, answer 3
```

- `d` to supply your own quiz duration, the value must be parsable into `time.Duration` format.

Example:

To set your quiz time to 10 seconds
```
go build . && ./timed-quiz -d 10s
```

To set your quiz time to 1 hours
```
go build . && ./timed-quiz -d 10s
```

## Future features
- [ ] Trim/clean to help ensure correct answers with extra whitespace/capitalization are not considered incorrect
- [ ] Shuffle quiz order
- [ ] Ability to restart the quiz
- [ ] Ability to show incorrect answer
