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
