-- +goose Up
-- +goose StatementBegin


CREATE TABLE objectives (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    type VARCHAR(255) NOT NULL CHECK (type IN ('year', 'month', 'week')),
    abandoned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE goals (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    abandoned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE objectives_goals (
    objective_id UUID NOT NULL,
    goal_id UUID NOT NULL,
    PRIMARY KEY (objective_id, goal_id),
    FOREIGN KEY (objective_id) REFERENCES objectives (id) ON DELETE CASCADE,
    FOREIGN KEY (goal_id) REFERENCES goals (id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE objectives_goals;
DROP TABLE goals;
DROP TABLE objectives;

-- +goose StatementEnd
