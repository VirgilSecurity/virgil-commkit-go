/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.9.4 at Wed Mar 24 14:44:04 2021. */

#ifndef PB_VSSQ_PB_VSSQ_CLOUDFILESYSTEM_PB_H_INCLUDED
#define PB_VSSQ_PB_VSSQ_CLOUDFILESYSTEM_PB_H_INCLUDED
#include <pb.h>

#include "google/protobuf/timestamp.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _vssq_pb_Permission {
    vssq_pb_Permission_PERMISSION_ADMIN = 0,
    vssq_pb_Permission_PERMISSION_USER = 1
} vssq_pb_Permission;
#define _vssq_pb_Permission_MIN vssq_pb_Permission_PERMISSION_ADMIN
#define _vssq_pb_Permission_MAX vssq_pb_Permission_PERMISSION_USER
#define _vssq_pb_Permission_ARRAYSIZE ((vssq_pb_Permission)(vssq_pb_Permission_PERMISSION_USER+1))

/* Struct definitions */
typedef struct _vssq_pb_CreateFileReq {
    char name[1025];
    char type[101];
    uint64_t size;
    char folder_id[33];
    pb_bytes_array_t *file_encrypted_key;
/* @@protoc_insertion_point(struct:vssq_pb_CreateFileReq) */
} vssq_pb_CreateFileReq;

typedef struct _vssq_pb_CreateFolderReq {
    char name[1025];
    char parent_folder_id[33];
    pb_bytes_array_t *folder_encrypted_key;
    pb_bytes_array_t *folder_public_key;
    pb_size_t users_count;
    struct _vssq_pb_User *users;
/* @@protoc_insertion_point(struct:vssq_pb_CreateFolderReq) */
} vssq_pb_CreateFolderReq;

typedef struct _vssq_pb_DeleteFileReq {
    char id[33];
/* @@protoc_insertion_point(struct:vssq_pb_DeleteFileReq) */
} vssq_pb_DeleteFileReq;

typedef struct _vssq_pb_DeleteFolderReq {
    char id[33];
/* @@protoc_insertion_point(struct:vssq_pb_DeleteFolderReq) */
} vssq_pb_DeleteFolderReq;

typedef struct _vssq_pb_File {
    char id[33];
    char name[1025];
    char type[101];
    uint64_t size;
    google_protobuf_Timestamp created_at;
    google_protobuf_Timestamp updated_at;
    char updated_by[257];
/* @@protoc_insertion_point(struct:vssq_pb_File) */
} vssq_pb_File;

typedef struct _vssq_pb_Folder {
    char id[33];
    char name[1025];
    google_protobuf_Timestamp created_at;
    google_protobuf_Timestamp updated_at;
    char updated_by[257];
    char shared_group_id[33];
/* @@protoc_insertion_point(struct:vssq_pb_Folder) */
} vssq_pb_Folder;

typedef struct _vssq_pb_GetFileLinkReq {
    char id[33];
/* @@protoc_insertion_point(struct:vssq_pb_GetFileLinkReq) */
} vssq_pb_GetFileLinkReq;

typedef struct _vssq_pb_GetFileLinkResp {
    char download_link[2049];
    pb_bytes_array_t *file_encrypted_key;
/* @@protoc_insertion_point(struct:vssq_pb_GetFileLinkResp) */
} vssq_pb_GetFileLinkResp;

typedef struct _vssq_pb_GetSharedGroupReq {
    char id[33];
/* @@protoc_insertion_point(struct:vssq_pb_GetSharedGroupReq) */
} vssq_pb_GetSharedGroupReq;

typedef struct _vssq_pb_Pagination {
    uint64_t limit;
    uint64_t offset;
/* @@protoc_insertion_point(struct:vssq_pb_Pagination) */
} vssq_pb_Pagination;

typedef struct _vssq_pb_SetSharedGroupReq {
    char id[33];
    pb_bytes_array_t *entry_encrypted_key;
    pb_size_t users_count;
    struct _vssq_pb_User *users;
/* @@protoc_insertion_point(struct:vssq_pb_SetSharedGroupReq) */
} vssq_pb_SetSharedGroupReq;

typedef struct _vssq_pb_SharedGroup {
    char id[33];
    pb_size_t users_count;
    struct _vssq_pb_User *users;
/* @@protoc_insertion_point(struct:vssq_pb_SharedGroup) */
} vssq_pb_SharedGroup;

typedef struct _vssq_pb_User {
    char identity[513];
    vssq_pb_Permission permission;
/* @@protoc_insertion_point(struct:vssq_pb_User) */
} vssq_pb_User;

typedef struct _vssq_pb_CreateFileResp {
    vssq_pb_File file;
    char upload_link[2049];
/* @@protoc_insertion_point(struct:vssq_pb_CreateFileResp) */
} vssq_pb_CreateFileResp;

typedef struct _vssq_pb_CreateFolderResp {
    vssq_pb_Folder folder;
/* @@protoc_insertion_point(struct:vssq_pb_CreateFolderResp) */
} vssq_pb_CreateFolderResp;

typedef struct _vssq_pb_GetSharedGroupResp {
    vssq_pb_SharedGroup shared_group;
/* @@protoc_insertion_point(struct:vssq_pb_GetSharedGroupResp) */
} vssq_pb_GetSharedGroupResp;

typedef struct _vssq_pb_ListFolderReq {
    char folder_id[33];
    vssq_pb_Pagination pagination;
/* @@protoc_insertion_point(struct:vssq_pb_ListFolderReq) */
} vssq_pb_ListFolderReq;

typedef struct _vssq_pb_ListFolderResp {
    uint64_t total_file_count;
    uint64_t total_folder_count;
    pb_size_t files_count;
    struct _vssq_pb_File *files;
    pb_size_t folders_count;
    struct _vssq_pb_Folder *folders;
    pb_bytes_array_t *folder_encrypted_key;
    pb_bytes_array_t *folder_public_key;
    vssq_pb_Folder current_folder;
/* @@protoc_insertion_point(struct:vssq_pb_ListFolderResp) */
} vssq_pb_ListFolderResp;

/* Default values for struct fields */

/* Initializer values for message structs */
#define vssq_pb_CreateFileReq_init_default       {"", "", 0, "", NULL}
#define vssq_pb_CreateFileResp_init_default      {vssq_pb_File_init_default, ""}
#define vssq_pb_GetFileLinkReq_init_default      {""}
#define vssq_pb_GetFileLinkResp_init_default     {"", NULL}
#define vssq_pb_DeleteFileReq_init_default       {""}
#define vssq_pb_CreateFolderReq_init_default     {"", "", NULL, NULL, 0, NULL}
#define vssq_pb_CreateFolderResp_init_default    {vssq_pb_Folder_init_default}
#define vssq_pb_ListFolderReq_init_default       {"", vssq_pb_Pagination_init_default}
#define vssq_pb_ListFolderResp_init_default      {0, 0, 0, NULL, 0, NULL, NULL, NULL, vssq_pb_Folder_init_default}
#define vssq_pb_DeleteFolderReq_init_default     {""}
#define vssq_pb_GetSharedGroupReq_init_default   {""}
#define vssq_pb_GetSharedGroupResp_init_default  {vssq_pb_SharedGroup_init_default}
#define vssq_pb_SetSharedGroupReq_init_default   {"", NULL, 0, NULL}
#define vssq_pb_File_init_default                {"", "", "", 0, google_protobuf_Timestamp_init_default, google_protobuf_Timestamp_init_default, ""}
#define vssq_pb_Folder_init_default              {"", "", google_protobuf_Timestamp_init_default, google_protobuf_Timestamp_init_default, "", ""}
#define vssq_pb_Pagination_init_default          {0, 0}
#define vssq_pb_User_init_default                {"", _vssq_pb_Permission_MIN}
#define vssq_pb_SharedGroup_init_default         {"", 0, NULL}
#define vssq_pb_CreateFileReq_init_zero          {"", "", 0, "", NULL}
#define vssq_pb_CreateFileResp_init_zero         {vssq_pb_File_init_zero, ""}
#define vssq_pb_GetFileLinkReq_init_zero         {""}
#define vssq_pb_GetFileLinkResp_init_zero        {"", NULL}
#define vssq_pb_DeleteFileReq_init_zero          {""}
#define vssq_pb_CreateFolderReq_init_zero        {"", "", NULL, NULL, 0, NULL}
#define vssq_pb_CreateFolderResp_init_zero       {vssq_pb_Folder_init_zero}
#define vssq_pb_ListFolderReq_init_zero          {"", vssq_pb_Pagination_init_zero}
#define vssq_pb_ListFolderResp_init_zero         {0, 0, 0, NULL, 0, NULL, NULL, NULL, vssq_pb_Folder_init_zero}
#define vssq_pb_DeleteFolderReq_init_zero        {""}
#define vssq_pb_GetSharedGroupReq_init_zero      {""}
#define vssq_pb_GetSharedGroupResp_init_zero     {vssq_pb_SharedGroup_init_zero}
#define vssq_pb_SetSharedGroupReq_init_zero      {"", NULL, 0, NULL}
#define vssq_pb_File_init_zero                   {"", "", "", 0, google_protobuf_Timestamp_init_zero, google_protobuf_Timestamp_init_zero, ""}
#define vssq_pb_Folder_init_zero                 {"", "", google_protobuf_Timestamp_init_zero, google_protobuf_Timestamp_init_zero, "", ""}
#define vssq_pb_Pagination_init_zero             {0, 0}
#define vssq_pb_User_init_zero                   {"", _vssq_pb_Permission_MIN}
#define vssq_pb_SharedGroup_init_zero            {"", 0, NULL}

/* Field tags (for use in manual encoding/decoding) */
#define vssq_pb_CreateFileReq_name_tag           1
#define vssq_pb_CreateFileReq_type_tag           2
#define vssq_pb_CreateFileReq_size_tag           3
#define vssq_pb_CreateFileReq_folder_id_tag      4
#define vssq_pb_CreateFileReq_file_encrypted_key_tag 5
#define vssq_pb_CreateFolderReq_name_tag         1
#define vssq_pb_CreateFolderReq_parent_folder_id_tag 2
#define vssq_pb_CreateFolderReq_folder_encrypted_key_tag 3
#define vssq_pb_CreateFolderReq_folder_public_key_tag 4
#define vssq_pb_CreateFolderReq_users_tag        5
#define vssq_pb_DeleteFileReq_id_tag             1
#define vssq_pb_DeleteFolderReq_id_tag           1
#define vssq_pb_File_id_tag                      1
#define vssq_pb_File_name_tag                    2
#define vssq_pb_File_type_tag                    3
#define vssq_pb_File_size_tag                    4
#define vssq_pb_File_created_at_tag              5
#define vssq_pb_File_updated_at_tag              6
#define vssq_pb_File_updated_by_tag              7
#define vssq_pb_Folder_id_tag                    1
#define vssq_pb_Folder_name_tag                  2
#define vssq_pb_Folder_created_at_tag            3
#define vssq_pb_Folder_updated_at_tag            4
#define vssq_pb_Folder_updated_by_tag            5
#define vssq_pb_Folder_shared_group_id_tag       6
#define vssq_pb_GetFileLinkReq_id_tag            1
#define vssq_pb_GetFileLinkResp_download_link_tag 1
#define vssq_pb_GetFileLinkResp_file_encrypted_key_tag 2
#define vssq_pb_GetSharedGroupReq_id_tag         1
#define vssq_pb_Pagination_limit_tag             1
#define vssq_pb_Pagination_offset_tag            2
#define vssq_pb_SetSharedGroupReq_id_tag         1
#define vssq_pb_SetSharedGroupReq_entry_encrypted_key_tag 2
#define vssq_pb_SetSharedGroupReq_users_tag      3
#define vssq_pb_SharedGroup_id_tag               1
#define vssq_pb_SharedGroup_users_tag            2
#define vssq_pb_User_identity_tag                1
#define vssq_pb_User_permission_tag              2
#define vssq_pb_CreateFileResp_file_tag          1
#define vssq_pb_CreateFileResp_upload_link_tag   2
#define vssq_pb_CreateFolderResp_folder_tag      1
#define vssq_pb_GetSharedGroupResp_shared_group_tag 1
#define vssq_pb_ListFolderReq_folder_id_tag      1
#define vssq_pb_ListFolderReq_pagination_tag     2
#define vssq_pb_ListFolderResp_total_file_count_tag 1
#define vssq_pb_ListFolderResp_total_folder_count_tag 2
#define vssq_pb_ListFolderResp_files_tag         3
#define vssq_pb_ListFolderResp_folders_tag       4
#define vssq_pb_ListFolderResp_folder_encrypted_key_tag 5
#define vssq_pb_ListFolderResp_folder_public_key_tag 6
#define vssq_pb_ListFolderResp_current_folder_tag 7

/* Struct field encoding specification for nanopb */
extern const pb_field_t vssq_pb_CreateFileReq_fields[6];
extern const pb_field_t vssq_pb_CreateFileResp_fields[3];
extern const pb_field_t vssq_pb_GetFileLinkReq_fields[2];
extern const pb_field_t vssq_pb_GetFileLinkResp_fields[3];
extern const pb_field_t vssq_pb_DeleteFileReq_fields[2];
extern const pb_field_t vssq_pb_CreateFolderReq_fields[6];
extern const pb_field_t vssq_pb_CreateFolderResp_fields[2];
extern const pb_field_t vssq_pb_ListFolderReq_fields[3];
extern const pb_field_t vssq_pb_ListFolderResp_fields[8];
extern const pb_field_t vssq_pb_DeleteFolderReq_fields[2];
extern const pb_field_t vssq_pb_GetSharedGroupReq_fields[2];
extern const pb_field_t vssq_pb_GetSharedGroupResp_fields[2];
extern const pb_field_t vssq_pb_SetSharedGroupReq_fields[4];
extern const pb_field_t vssq_pb_File_fields[8];
extern const pb_field_t vssq_pb_Folder_fields[7];
extern const pb_field_t vssq_pb_Pagination_fields[3];
extern const pb_field_t vssq_pb_User_fields[3];
extern const pb_field_t vssq_pb_SharedGroup_fields[3];

/* Maximum encoded size of messages (where known) */
/* vssq_pb_CreateFileReq_size depends on runtime parameters */
#define vssq_pb_CreateFileResp_size              3540
#define vssq_pb_GetFileLinkReq_size              35
/* vssq_pb_GetFileLinkResp_size depends on runtime parameters */
#define vssq_pb_DeleteFileReq_size               35
/* vssq_pb_CreateFolderReq_size depends on runtime parameters */
#define vssq_pb_CreateFolderResp_size            1409
#define vssq_pb_ListFolderReq_size               59
/* vssq_pb_ListFolderResp_size depends on runtime parameters */
#define vssq_pb_DeleteFolderReq_size             35
#define vssq_pb_GetSharedGroupReq_size           35
/* vssq_pb_GetSharedGroupResp_size depends on runtime parameters */
/* vssq_pb_SetSharedGroupReq_size depends on runtime parameters */
#define vssq_pb_File_size                        1485
#define vssq_pb_Folder_size                      1406
#define vssq_pb_Pagination_size                  22
#define vssq_pb_User_size                        518
/* vssq_pb_SharedGroup_size depends on runtime parameters */

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define VSSQ_CLOUDFILESYSTEM_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
