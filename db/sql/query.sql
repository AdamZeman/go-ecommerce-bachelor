-- name: GetProducts :many
SELECT * from products
ORDER BY id;

-- name: GetProductsByCategoryId :many
SELECT p.*
FROM products p
         JOIN product_categories pc ON p.id = pc.product_id
WHERE pc.category_id = $1;

-- name: GetCategoriesFilteredProductID :many
SELECT
    c.*,
    CASE
        WHEN pc.product_id IS NOT NULL THEN TRUE
        ELSE FALSE
        END AS has_product
FROM categories c
         LEFT JOIN product_categories pc
                   ON c.id = pc.category_id
                       AND pc.product_id = $1;

-- name: GetCategoryByID :one
SELECT * from categories
WHERE id = $1;


-- name: GetCategories :many
SELECT * from categories
order by id asc ;

-- name: GetProductById :one
SELECT * from products where id = $1;

-- name: GetProductVariantsByProductId :many
SELECT pv.*
FROM product_variants pv
         JOIN products p ON pv.product_id = p.id
WHERE p.id = $1;

-- name: GetBasketItemsByUserId :many
SELECT *
FROM basket_items bi
         JOIN product_variants pv on pv.id = bi.product_variant_id
         JOIN products p on p.id = pv.product_id
WHERE bi.user_id = $1;

-- name: GetBasketItemsByUserEmail :many
SELECT *
FROM basket_items bi
         JOIN product_variants pv on pv.id = bi.product_variant_id
         JOIN products p on p.id = pv.product_id
         JOIN users u on u.id = bi.user_id
WHERE u.email = $1;


-- name: InsertBasketItem :one
INSERT INTO basket_items (user_id, product_variant_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: RemoveFromBasket :one
DELETE FROM basket_items
WHERE id = $1
RETURNING *;

-- name: UpsertUser :one
INSERT INTO users (google_id, email, name, avatar_url)
VALUES ($1, $2, $3, $4)
ON CONFLICT (google_id)
    DO UPDATE SET
                  email = EXCLUDED.email,
                  name = EXCLUDED.name,
                  avatar_url = EXCLUDED.avatar_url
RETURNING *;

-- name: InsertMessage :one
INSERT INTO conversation_messages (conversation_id, sender_id, content)
VALUES ($1, (SELECT id FROM users WHERE email = $2), $3)
RETURNING *;

-- name: GetMessages :many
SELECT cm.*, u.email
FROM conversation_messages cm
         JOIN users u on cm.sender_id = u.id
ORDER BY cm.created_at ASC;

-- name: GetMessagesByConversationId :many
SELECT cm.*, u.email
FROM conversation_messages cm
         JOIN conversations c ON cm.conversation_id = c.id
         JOIN users u ON cm.sender_id = u.id
WHERE c.id = $1
ORDER BY cm.created_at ASC;

-- name: GetConversations :many
SELECT c.*
FROM conversations c;

-- name: GetConversationsByUserId :many
SELECT DISTINCT c.*
FROM conversations c
         LEFT JOIN conversations_users cu ON cu.conversation_id = c.id
         LEFT JOIN users u ON u.id = cu.user_id
WHERE u.email = $1;


-- name: InsertConversation :one
INSERT INTO conversations (name, order_id)
VALUES ($1, $2)
RETURNING *;

-- name: InsertConversationUser :one
INSERT INTO conversations_users (conversation_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetUsers :many
SELECT *
FROM users;

-- name: GetOrderItemsByUserId :many
SELECT *
FROM orders o
         JOIN order_items oi on oi.order_id = o.id
         JOIN product_variants pv on pv.id = oi.product_variant_id
         JOIN products p on p.id = pv.product_id
WHERE o.user_id = $1;

-- name: GetOrdersByUserId :many
SELECT *
FROM orders o
WHERE o.user_id = $1
ORDER BY o.created_at desc;

-- name: GetProductsByCategoriesAndName :many
SELECT p.*, s.*, (fi.product_id IS NOT NULL)::boolean AS "isFavourite"
FROM products p
         LEFT JOIN special s ON p.special = s.id
         LEFT JOIN favourite_items fi
                   ON fi.product_id = p.id AND fi.user_id = @user_id

WHERE p.name ILIKE '%' || @name_filter || '%'
  AND (
    array_length(@category_ids::int[], 1) IS NULL
        OR EXISTS (
        SELECT 1 FROM product_categories pc
        WHERE pc.product_id = p.id AND pc.category_id = ANY(@category_ids::int[])
    )
    )
  AND (
    array_length(@special_ids::int[], 1) IS NULL
        OR p.special = ANY(@special_ids::int[])
    )
  AND (
    (p.price >= @price_from::int)
        AND
    (p.price <= @price_to::int)
    )
ORDER BY p.id
LIMIT @limitVar::int
OFFSET @offsetVar::int;

-- name: GetFavouritesById :many
SELECT *
FROM favourite_items fi
         JOIN users u ON fi.user_id = u.id
         JOIN products p on p.id = fi.product_id
WHERE u.email = $1;

-- name: ProductVariantsByProductId :many
SELECT pv.*
FROM product_variants pv
WHERE pv.product_id = $1;

-- name: GetVariantsFilledProducts :many
SELECT *
FROM product_variants pv
         JOin products p on pv.product_id = p.id
order by pv.id asc ;

-- name: GetVariantByOptions :one
SELECT pv.*
FROM product_variants pv
WHERE product_id = $1
  AND option_1_value = $2
  AND option_2_value = $3;

-- name: InsertOrder :one
INSERT INTO orders (user_id, total_price, shipping_id)
Values ($1, $2, $3)
RETURNING *;

-- name: InsertOrderItem :one
INSERT INTO order_items (order_id, product_variant_id, quantity)
Values ($1, $2, $3)
RETURNING *;

-- name: DeleteItemsFromBasket :exec
DELETE FROM basket_items bi
WHERE bi.id = $1;

-- name: InsertItemToFavourites :one
INSERT INTO favourite_items (user_id, product_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveItemFromFavourites :exec
DELETE FROM favourite_items fi
WHERE fi.product_id = $1 and fi.user_id = $2;

-- name: GetFavouriteByProductId :many
SELECT *
FROM favourite_items fi
WHERE fi.product_id = $1 and fi.user_id = $2;

-- name: GetOpenConversationById :many
SELECT *
FROM conversations c
         JOIN orders o on c.order_id = o.id
WHERE c.order_id = $1 AND o.status != $2;

-- name: InsertProduct :exec
INSERT INTO products (name, description, price, option_1_name, option_2_name, special)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateProduct :exec
UPDATE products
SET
    name = $2,
    description = $3,
    price = $4,
    option_1_name = $5,
    option_2_name = $6,
    special = $7
WHERE id = $1;

-- name: DeleteProductByID :exec
DELETE FROM products
WHERE id = $1;

-- name: AddProductCategoryByIDs :exec
INSERT INTO product_categories (product_id, category_id)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT 1 FROM product_categories
    WHERE product_id = $1 AND category_id = $2
);

-- name: DeleteProductCategoryByIDs :exec
DELETE FROM product_categories
WHERE product_id=$1 and category_id=$2;

-- name: GetVariantById :one
Select *
From product_variants
WHERE id = $1;

-- name: UpdateVariant :exec
UPDATE product_variants
SET
    sku = $2,
    stock_quantity = $3,
    price = $4,
    option_1_value = $5,
    option_2_value = $6
WHERE id = $1;

-- name: InsertVariant :exec
INSERT INTO product_variants (product_id, sku, price, stock_quantity, option_1_value, option_2_value)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: DeleteVariantByID :exec
DELETE FROM product_variants
WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories
SET
    name = $2
WHERE id = $1;

-- name: InsertCategory :exec
INSERT INTO categories (name)
VALUES ($1);

-- name: DeleteCategoryByID :exec
DELETE FROM categories
WHERE id = $1;

-- name: GetOrdersFillUser :many
SELECT *
From orders o
         JOIN users on o.user_id = users.id
order by o.id asc ;

-- name: GetOrdersFillUserByStatus :many
SELECT *
From orders o
         JOIN users on o.user_id = users.id
WHERE status = $1
order by o.id asc ;


-- name: GetOrdersFillUserConvByStatus :many
SELECT o.*, u.email, c.id as convId
From orders o
         JOIN users u on o.user_id = u.id
         JOIN conversations c on o.id = c.order_id
WHERE status = $1
order by o.id asc ;


-- name: GetOrderItemsByOrderID :many
SELECT *
FROM orders o
         JOIN order_items oi on oi.order_id = o.id
         JOIN product_variants pv on pv.id = oi.product_variant_id
         JOIN products p on p.id = pv.product_id
WHERE o.id = $1;

-- name: GetOrderByID :one
Select *
From orders
WHERE id = $1;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: CountConvByUserConv :one
SELECT count(*)
FROM conversations_users
WHERE user_id = $1 AND conversation_id = $2;

-- name: GetReviewsFillUserByProductID :many
SELECT *
FROM reviews r
         JOIN users u on r.sender_id = u.id
WHERE r.product_id = $1;

-- name: GetCategoriesByProductID :many
SELECT c.*
From categories c
         JOIN product_categories pc on c.id = pc.category_id
WHERE pc.product_id = $1;

-- name: GetSpecial :many
Select *
FROM special;

-- name: UpdateSpecial :exec
UPDATE special
SET
    name = $2
WHERE id = $1;

-- name: InsertSpecial :exec
INSERT INTO special (name)
VALUES ($1);

-- name: DeleteSpecialByID :exec
DELETE FROM special
WHERE id = $1;

-- name: GetSpecialByID :one
SELECT * from special
WHERE id = $1;

-- name: GetProductCategoriesManyToMany :many
SELECT * from product_categories;

-- name: GetBasketItemByBasketItemId :one
SELECT *
FROM basket_items bi
         JOIN product_variants pv on pv.id = bi.product_variant_id
         JOIN products p on p.id = pv.product_id
WHERE bi.id = $1;

-- name: InsertReview :one
INSERT INTO reviews
(sender_id, product_id, rating, content)
VALUES ($1, $2, $3, $4)
returning *;

-- name: InsertShipping :one
INSERT INTO shippings
(user_id, phone_number, address, house_number, postal_code, city, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetLatestShippingByUserId :one
SELECT sh.*
FROM shippings sh
         JOIN users u on sh.user_id = u.id
WHERE u.email = $1
ORDER BY sh.id desc
LIMIT 1;