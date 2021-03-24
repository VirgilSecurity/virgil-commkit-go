/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.9.4 at Wed Mar 24 13:44:35 2021. */

#ifndef PB_UOKMSMODELS_PB_H_INCLUDED
#define PB_UOKMSMODELS_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _UOKMSProofOfSuccess {
    pb_byte_t term1[65];
    pb_byte_t term2[65];
    pb_byte_t blind_x[32];
/* @@protoc_insertion_point(struct:UOKMSProofOfSuccess) */
} UOKMSProofOfSuccess;

typedef struct _DecryptResponse {
    pb_byte_t v[65];
    UOKMSProofOfSuccess proof;
/* @@protoc_insertion_point(struct:DecryptResponse) */
} DecryptResponse;

/* Default values for struct fields */

/* Initializer values for message structs */
#define UOKMSProofOfSuccess_init_default         {{0}, {0}, {0}}
#define DecryptResponse_init_default             {{0}, UOKMSProofOfSuccess_init_default}
#define UOKMSProofOfSuccess_init_zero            {{0}, {0}, {0}}
#define DecryptResponse_init_zero                {{0}, UOKMSProofOfSuccess_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define UOKMSProofOfSuccess_term1_tag            1
#define UOKMSProofOfSuccess_term2_tag            2
#define UOKMSProofOfSuccess_blind_x_tag          3
#define DecryptResponse_v_tag                    1
#define DecryptResponse_proof_tag                2

/* Struct field encoding specification for nanopb */
extern const pb_field_t UOKMSProofOfSuccess_fields[4];
extern const pb_field_t DecryptResponse_fields[3];

/* Maximum encoded size of messages (where known) */
#define UOKMSProofOfSuccess_size                 168
#define DecryptResponse_size                     238

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define UOKMSMODELS_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
