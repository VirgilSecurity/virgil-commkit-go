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
//  This file contains platform specific information that is known during compilation.
// --------------------------------------------------------------------------

#ifndef VSSK_PLATFORM_H_INCLUDED
#define VSSK_PLATFORM_H_INCLUDED

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

#ifndef VSSK_HAVE_ASSERT_H
#define VSSK_HAVE_ASSERT_H 1
#endif

#ifndef VSSK_HAVE_STDATOMIC_H
#define VSSK_HAVE_STDATOMIC_H 1
#endif

#ifndef VSSK_SHARED_LIBRARY
#define VSSK_SHARED_LIBRARY 0
#endif

#ifndef VSSK_MULTI_THREADING
#define VSSK_MULTI_THREADING 1
#endif

#ifndef VSSK_ERROR
#define VSSK_ERROR 1
#endif

#ifndef VSSK_KEYKNOX_CLIENT
#define VSSK_KEYKNOX_CLIENT 1
#endif

#ifndef VSSK_KEYKNOX_ENTRY
#define VSSK_KEYKNOX_ENTRY 1
#endif

//
//  Defines namespace include prefix for project 'common'.
//
#if !defined(VSSK_INTERNAL_BUILD)
#define VSSK_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK 0
#else
#define VSSK_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK 0
#endif

//
//  Defines namespace include prefix for project 'foundation'.
//
#if !defined(VSSK_INTERNAL_BUILD)
#define VSSK_IMPORT_PROJECT_FOUNDATION_FROM_FRAMEWORK 0
#else
#define VSSK_IMPORT_PROJECT_FOUNDATION_FROM_FRAMEWORK 0
#endif

//
//  Defines namespace include prefix for project 'core sdk'.
//
#if !defined(VSSK_INTERNAL_BUILD)
#define VSSK_IMPORT_PROJECT_CORE_SDK_FROM_FRAMEWORK 0
#else
#define VSSK_IMPORT_PROJECT_CORE_SDK_FROM_FRAMEWORK 0
#endif


// --------------------------------------------------------------------------
//  Generated section end.
// clang-format on
// --------------------------------------------------------------------------
//  @end


#ifdef __cplusplus
}
#endif


//  @footer
#endif // VSSK_PLATFORM_H_INCLUDED
//  @end
