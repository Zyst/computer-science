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
select name, rID, mID
from Rating, Reviewer using (rID)
order by mID
