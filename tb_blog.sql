CREATE TABLE blog
(
    id         SERIAL PRIMARY KEY,                    -- 唯一标识博客的ID，自动递增
    title      VARCHAR(255) NOT NULL,                 -- 博客文章的标题，最大255字符
    content    TEXT         NOT NULL,                 -- 博客文章的内容，使用TEXT存储长文本
    created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP, -- 创建时间，默认为当前时间
    updated_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP, -- 更新时间，默认为当前时间
    category   VARCHAR(100),                          -- 博客分类
    tags       VARCHAR(255),                          -- 文章的标签，可以用逗号分隔多个标签
    view_count INT         DEFAULT 0                  -- 文章的查看次数，默认为0
);


-- 插入第一篇博客文章
INSERT INTO blog (title, content, category, tags, view_count)
VALUES
    ('Introduction to PostgreSQL',
     'PostgreSQL is a powerful, open source object-relational database system.',
     'Database',
     'postgresql,database,sql',
     120);

-- 插入第二篇博客文章
INSERT INTO blog (title, content, category, tags, view_count)
VALUES
    ('Building a Simple REST API with Node.js',
     'In this tutorial, we will learn how to build a simple REST API using Node.js and Express.',
     'Web Development',
     'node.js,express,api',
     200);

-- 插入第三篇博客文章
INSERT INTO blog (title, content, category, tags, view_count)
VALUES
    ('Understanding CSS Grid Layout',
     'CSS Grid Layout is a two-dimensional layout system for the web. It lets you layout items in rows and columns.',
     'Web Design',
     'css,grid,layout',
     80);

-- 插入第四篇博客文章
INSERT INTO blog (title, content, category, tags, view_count)
VALUES
    ('Getting Started with React Hooks',
     'React Hooks allow you to use state and other React features without writing a class.',
     'Frontend',
     'react,hooks,javascript',
     150);

-- 插入第五篇博客文章
INSERT INTO blog (title, content, category, tags, view_count)
VALUES
    ('How to Secure Your Web Application',
     'In this guide, we will cover best practices for securing web applications.',
     'Security',
     'security,web,application',
     95);