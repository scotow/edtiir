# edtiir 📆

A Google Spreadsheet wrapper poiting to the current week of the TIIR timetable (Lille University CS Master Degree).

There is two commands in the [`cmd`](https://github.com/Scotow/edtiir/tree/master/cmd) directory:

- [`link`](https://github.com/Scotow/edtiir/tree/master/cmd/link) prints the link of the current week to `stdout`.
- [`web`](https://github.com/Scotow/edtiir/tree/master/cmd/web) starts a web server on port 8080 and return a HTTP 307 response pointing to the Google Spreadsheet link and the current week.

NB: This project was made to save us 2s everytime my friends or I wanted to see our next courses and should not be considered as a serious project.
