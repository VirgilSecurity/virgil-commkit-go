//  @license
// --------------------------------------------------------------------------
//  Copyright (C) 2015-2020 Virgil Security, Inc.
//
//  All rights reserved.
//
//  Redistribution and use in source and binary forms, with or without
//  modification, are permitted provided that the following conditions are
//  met:
//
//      (1) Redistributions of source code must retain the above copyright
//      notice, this list of conditions and the following disclaimer.
//
//      (2) Redistributions in binary form must reproduce the above copyright
//      notice, this list of conditions and the following disclaimer in
//      the documentation and/or other materials provided with the
//      distribution.
//
//      (3) Neither the name of the copyright holder nor the names of its
//      contributors may be used to endorse or promote products derived from
//      this software without specific prior written permission.
//
//  THIS SOFTWARE IS PROVIDED BY THE AUTHOR ''AS IS'' AND ANY EXPRESS OR
//  IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//  DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
//  INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
//  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
//  HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
//  STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
//  IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
//  POSSIBILITY OF SUCH DAMAGE.
//
//  Lead Maintainer: Virgil Security Inc. <support@virgilsecurity.com>
// --------------------------------------------------------------------------
// clang-format off


//  @warning
// --------------------------------------------------------------------------
//  This file is partially generated.
//  Generated blocks are enclosed between tags [@<tag>, @end].
//  User's code can be added between tags [@end, @<tag>].
// --------------------------------------------------------------------------


//  @description
// --------------------------------------------------------------------------
//  Handles a list of "card" class objects.
// --------------------------------------------------------------------------

#ifndef VSSC_CARD_LIST_H_INCLUDED
#define VSSC_CARD_LIST_H_INCLUDED

#include "vssc_library.h"
#include "vssc_card.h"
#include "vssc_error.h"

#if !VSSC_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK
#   include <virgil/crypto/common/vsc_str.h>
#endif

#if VSSC_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK
#   include <VSCCommon/vsc_str.h>
#endif

// clang-format on
//  @end


#ifdef __cplusplus
extern "C" {
#endif


//  @generated
// --------------------------------------------------------------------------
// clang-format off
//  Generated section start.
// --------------------------------------------------------------------------

//
//  Handle 'card list' context.
//
#ifndef VSSC_CARD_LIST_T_DEFINED
#define VSSC_CARD_LIST_T_DEFINED
    typedef struct vssc_card_list_t vssc_card_list_t;
#endif // VSSC_CARD_LIST_T_DEFINED

//
//  Return size of 'vssc_card_list_t'.
//
VSSC_PUBLIC size_t
vssc_card_list_ctx_size(void);

//
//  Perform initialization of pre-allocated context.
//
VSSC_PUBLIC void
vssc_card_list_init(vssc_card_list_t *self);

//
//  Release all inner resources including class dependencies.
//
VSSC_PUBLIC void
vssc_card_list_cleanup(vssc_card_list_t *self);

//
//  Allocate context and perform it's initialization.
//
VSSC_PUBLIC vssc_card_list_t *
vssc_card_list_new(void);

//
//  Release all inner resources and deallocate context if needed.
//  It is safe to call this method even if the context was statically allocated.
//
VSSC_PUBLIC void
vssc_card_list_delete(const vssc_card_list_t *self);

//
//  Delete given context and nullifies reference.
//  This is a reverse action of the function 'vssc_card_list_new ()'.
//
VSSC_PUBLIC void
vssc_card_list_destroy(vssc_card_list_t **self_ref);

//
//  Copy given class context by increasing reference counter.
//
VSSC_PUBLIC vssc_card_list_t *
vssc_card_list_shallow_copy(vssc_card_list_t *self);

//
//  Copy given class context by increasing reference counter.
//  Reference counter is internally synchronized, so constness is presumed.
//
VSSC_PUBLIC const vssc_card_list_t *
vssc_card_list_shallow_copy_const(const vssc_card_list_t *self);

//
//  Remove current node.
//
VSSC_PRIVATE void
vssc_card_list_remove_self(vssc_card_list_t *self);

//
//  Return true if given list has item.
//
VSSC_PUBLIC bool
vssc_card_list_has_item(const vssc_card_list_t *self);

//
//  Return list item.
//
VSSC_PUBLIC const vssc_card_t *
vssc_card_list_item(const vssc_card_list_t *self);

//
//  Return true if list has next item.
//
VSSC_PUBLIC bool
vssc_card_list_has_next(const vssc_card_list_t *self);

//
//  Return next list node if exists, or NULL otherwise.
//
VSSC_PUBLIC const vssc_card_list_t *
vssc_card_list_next(const vssc_card_list_t *self);

//
//  Return true if list has previous item.
//
VSSC_PUBLIC bool
vssc_card_list_has_prev(const vssc_card_list_t *self);

//
//  Return previous list node if exists, or NULL otherwise.
//
VSSC_PUBLIC const vssc_card_list_t *
vssc_card_list_prev(const vssc_card_list_t *self);

//
//  Remove all items.
//
VSSC_PUBLIC void
vssc_card_list_clear(vssc_card_list_t *self);

//
//  Find first card with it's identity.
//
VSSC_PUBLIC const vssc_card_t *
vssc_card_list_find_with_identity(const vssc_card_list_t *self, vsc_str_t identity, vssc_error_t *error);

//
//  Find card with it's identifier.
//
VSSC_PUBLIC const vssc_card_t *
vssc_card_list_find_with_identifier(const vssc_card_list_t *self, vsc_str_t identifier, vssc_error_t *error);


// --------------------------------------------------------------------------
//  Generated section end.
// clang-format on
// --------------------------------------------------------------------------
//  @end


#ifdef __cplusplus
}
#endif


//  @footer
#endif // VSSC_CARD_LIST_H_INCLUDED
//  @end
