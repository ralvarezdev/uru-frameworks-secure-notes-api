-- Drop all functions in the database
DROP FUNCTION IF EXISTS get_user_refresh_token_by_id;
DROP FUNCTION IF EXISTS list_user_refresh_tokens;
DROP FUNCTION IF EXISTS list_user_tokens;
DROP FUNCTION IF EXISTS list_user_tags;
DROP FUNCTION IF EXISTS list_user_note_versions;
DROP FUNCTION IF EXISTS list_user_note_tags;
DROP FUNCTION IF EXISTS sync_user_note_versions_by_last_synced_at;
DROP FUNCTION IF EXISTS sync_user_notes_by_last_synced_at;
DROP FUNCTION IF EXISTS sync_user_note_tags_by_last_synced_at;
DROP FUNCTION IF EXISTS sync_user_tags_by_last_synced_at;