-- Q1 Find the names of all reviewers who rated Gone with the Wind.
select distinct name
from Rating, Reviewer using (rID), Movie using (mID)
where title = "Gone with the Wind"

-- Q2 For any rating where the reviewer is the same as the director of the movie, return the reviewer name, movie title, and number of stars.
select name, title, stars
from Rating, Reviewer using (rID), Movie using (mID)
where director = name

-- Q3 Return all reviewer names and movie names together in a single list, alphabetized. (Sorting by the first name of the reviewer and first word in the title is fine; no need for special processing on last names or removing "The".)
select name
from Reviewer
union
select title
from Movie

-- Q4 Find the titles of all movies not reviewed by Chris Jackson
select title
from Movie
where Movie.mID not in (
  select mID
  from Rating, Reviewer using (rID)
  where name = "Chris Jackson")

-- Q5 For all pairs of reviewers such that both reviewers gave a rating to the same movie, return the names of both reviewers. Eliminate duplicates, don't pair reviewers with themselves, and include each pair only once. For each pair, return the names in the pair in alphabetical order.
select distinct
  (
    select name
    from Reviewer
    where rID = R1.rID
  ) nameR1,
  (
    select name
    from Reviewer
    where rID = R2.rID
  ) nameR2
from Rating as R1, Rating as R2
where R1.mID = R2.mID and R1.rID > R2.rID
order by nameR1, nameR2

-- Q6 For each rating that is the lowest (fewest stars) currently in the database, return the reviewer name, movie title, and number of stars.
select name, title, stars
from Rating, Movie using (mID), Reviewer using (rID)
where stars = (select min(stars) from Rating)

-- Q7 List movie titles and average ratings, from highest-rated to lowest-rated. If two or more movies have the same average rating, list them in alphabetical order.
select title, avg(stars) as avgStars
from Rating, Movie using (mID)
group by title
order by avgStars desc, title

-- Q8 Find the names of all reviewers who have contributed three or more ratings. (As an extra challenge, try writing the query without HAVING or without COUNT.)
select name
from Rating, Reviewer using (rID)
group by rID
having count(rID) > 2
