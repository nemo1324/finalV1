-- name: UpdateUserStatus :exec
UPDATE users
SET status = @status
WHERE id = @id;

-- name: UpdateUserName :exec
UPDATE users
SET name = @name
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id;

-- name: Logout :exec
UPDATE users
SET status = 'offline'
WHERE id = @id;
