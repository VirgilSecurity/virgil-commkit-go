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

#ifndef VSSB_PLATFORM_H_INCLUDED
#define VSSB_PLATFORM_H_INCLUDED

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

#ifndef VSSB_HAVE_ASSERT_H
#define VSSB_HAVE_ASSERT_H 1
#endif

#ifndef VSSB_HAVE_STDATOMIC_H
#define VSSB_HAVE_STDATOMIC_H 1
#endif

#ifndef VSSB_SHARED_LIBRARY
#define VSSB_SHARED_LIBRARY 0
#endif

#ifndef VSSB_MULTI_THREADING
#define VSSB_MULTI_THREADING 1
#endif

#ifndef VSSB_ERROR
#define VSSB_ERROR 1
#endif

#ifndef VSSB_ERROR_MESSAGE
#define VSSB_ERROR_MESSAGE 1
#endif

#ifndef VSSB_BRAINKEY_CLIENT
#define VSSB_BRAINKEY_CLIENT 1
#endif

#ifndef VSSB_BRAINKEY_HARDENED_POINT
#define VSSB_BRAINKEY_HARDENED_POINT 1
#endif

//
//  Defines namespace include prefix for project 'common'.
//
#if !defined(VSSB_INTERNAL_BUILD)
#define VSSB_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK 0
#else
#define VSSB_IMPORT_PROJECT_COMMON_FROM_FRAMEWORK 0
#endif

//
//  Defines namespace include prefix for project 'foundation'.
//
#if !defined(VSSB_INTERNAL_BUILD)
#define VSSB_IMPORT_PROJECT_FOUNDATION_FROM_FRAMEWORK 0
#else
#define VSSB_IMPORT_PROJECT_FOUNDATION_FROM_FRAMEWORK 0
#endif

//
//  Defines namespace include prefix for project 'core sdk'.
//
#if !defined(VSSB_INTERNAL_BUILD)
#define VSSB_IMPORT_PROJECT_CORE_SDK_FROM_FRAMEWORK 0
#else
#define VSSB_IMPORT_PROJECT_CORE_SDK_FROM_FRAMEWORK 0
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
#endif // VSSB_PLATFORM_H_INCLUDED
//  @end
