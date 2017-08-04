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
