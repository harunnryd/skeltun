# `/migration`

Put your `migration` files in here.

Shortcut.
-- Add a spatial index
DO
$$
BEGIN
    IF to_regclass('workspace_areas_gix') IS NULL THEN
      CREATE INDEX workspace_areas_gix ON workspace_areas USING GIST (geom);
    END IF;
END
$$;
