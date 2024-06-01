ALTER TABLE `todo` RENAME COLUMN `createdAt` TO `created_at`;
ALTER TABLE `todo` RENAME COLUMN `updatedAt` TO `updated_at`;

ALTER TABLE `todo_tag` RENAME COLUMN `createdAt` TO `created_at`;

ALTER TABLE `todosongroups` RENAME COLUMN `assignedAt` TO `assigned_at`;