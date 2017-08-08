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
