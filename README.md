# (D)ay (O)ne (D)owntime CLI - dod

I don't know about you, but I love [Day One](https://dayoneapp.com). It's absolutely changed my life.

I'm also obsessed with tracking the things I do in my downtime - books, games, TV, etc.

Day One is awesome at tracking these activities, and I love the way it displays the media I collect for each of the activities that I log in my journal.

![Dayone media display](./images/dayone-media-display.png)

However, come the end of the year I'd like to have a summary of all my downtime activity for the year - logged as another journal entry of course!

That's where this tool comes in.

## How it works

- Unfortunately, there is no API for Day One, so this tool relies on a `JSON` format export from Dayone to retrieve it's information.
- The tool relies on entries being well-tagged. I tag my entries based on:
  - `type`: books, videogames, movies, etc
  - `genre`: Theres no limit to how many genres, but they should be consistent in their usage
  - `rating`: #1star - #5stars. Storing rating as a tag allows for easier filtering and locating entries with the different ratings
  - `favorite`: If I really loved something, I tag it as such so it's easier to find later
  - `dnf` (did not finish): things so bad I just had to give up - what were they?

<WIP>

## TODO

- [ ] Process ratings (1-5 star) and report on average rating, top rated, bottom rated, etc
- [ ]
