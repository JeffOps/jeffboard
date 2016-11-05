SELECT f.* FROM (
  SELECT subject, text, thread_id, date_posted,
  rank() OVER (PARTITION BY thread_id ORDER BY date_posted DESC) AS drank, -- last post in thread has drank = 1
  rank() OVER (PARTITION BY thread_id ORDER BY date_posted ASC) AS arank,  -- first post in thread has arank = 1
  max(date_posted) OVER (PARTITION BY thread_id) AS last_post_in_thread FROM post
) AS f 
WHERE drank < 3 -- last two posts in each thread
OR arank = 1 -- first post
ORDER BY last_post_in_thread DESC, thread_id, arank, drank DESC;
