# reps

## Overview 

reps is a basic spaced repetition app.

reps includes a chrome extension that allows users to schedule any text for 
spaced repetition. Users are notified of the text repeatedly, with intervals 
between successive reminders being determined by exponential backoff.

Users are notified of their chosen text via SMS, to the phone number they provide.

## How it was built

reps was built to learn and experiment with some new tools:

- A Go Gin HTTP server, with Gorm for Postgres ORM support
- Temporal, a workflow orchestration framework, responsible for executing the 
long-living repetition workflows, with built-in workflow durability and retry behaviour
- The Twilio SMS framework for sending repetition texts, defined as a Temporal Activity.
