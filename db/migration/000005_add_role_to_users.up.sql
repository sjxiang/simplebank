
-- 非空约束，默认值 '储户'
ALTER TABLE "users" ADD COLUMN "role" varchar NOT NULL DEFAULT 'depositor';