ALTER TABLE tasks 
ADD COLUMN IF NOT EXISTS frequency JSONB DEFAULT NULL;

CREATE INDEX IF NOT EXISTS idx_tasks_frequency_type 
ON tasks ((frequency->>'type'));
