-- Q1 Find the names of all reviewers who rated Gone with the Wind.
select distinct name
from Rating, Reviewer using (rID), Movie using (mID)
where title = "Gone with the Wind"
