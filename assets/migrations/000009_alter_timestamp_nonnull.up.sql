ALTER TABLE `goadmin_menu`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_menu`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_operation_log`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_operation_log`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_permissions`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_permissions`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_role_menu`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_role_menu`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_role_permissions`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_role_permissions`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_role_users`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_role_users`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_roles`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_roles`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_session`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_session`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_site`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_site`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_user_permissions`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_user_permissions`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `goadmin_users`
    MODIFY `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `goadmin_users`
    MODIFY `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;