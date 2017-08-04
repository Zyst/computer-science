-- Q1 Find the titles of all movies directed by Steven Spielberg.
select title
from Movie
where director = "Steven Spielberg"

-- Q2 Find all years that have a movie that received a rating of 4 or 5, and sort them in increasing order
select distinct year
from Rating, Movie using(mID)
where stars = 5 or stars = 4
order by year

-- Q3 Find the titles of all movies that have no ratings.
select title
from Movie left join Rating on Movie.mID = Rating.mID
where rID is null

-- Q4 Some reviewers didn't provide a date with their rating. Find the names of all reviewers who have ratings with a NULL value for the date.
select name
from Rating, Reviewer using (rID)
where ratingDate is null

-- Q5 Write a query to return the ratings data in a more readable format: reviewer name, movie title, stars, and ratingDate. Also, sort the data, first by reviewer name, then by movie title, and lastly by number of stars.
select name, title, stars, ratingDate
from Movie, Reviewer, Rating using (mID, rID)
order by name, title, stars

-- Q6 For all cases where the same reviewer rated the same movie twice and gave it a higher rating the second time, return the reviewer's name and the title of the movie.
-- I'm like 100% sure there must be a way more elegant way to do this
select name, title
from Rating as R1, Rating as R2, Reviewer using (rID), Movie using (mID)
where R1.rID = R2.rID and R1.mID = R2.mID and R1.stars > R2.stars and R1.ratingDate > R2.ratingDate

-- Q7 For each movie that has at least one rating, find the highest number of stars that movie received. Return the movie title and number of stars. Sort by movie title.
select title, max(stars)
from Movie, Rating using (mID)
group by title
