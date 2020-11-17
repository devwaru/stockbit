SELECT u.id, u.user_name, p.user_name AS parent_user_name
FROM users u
LEFT JOIN users p NO p.id = u.parent;