package sdk_comm_kit

// #include <virgil/sdk/comm-kit/vssq_comm_kit_public.h>
import "C"
import "unsafe"
import "runtime"

/*
* Handles info required to download and decrypt file.
 */
type MessengerCloudFsFileDownloadInfo struct {
	cCtx *C.vssq_messenger_cloud_fs_file_download_info_t /*ct2*/
}

/* Handle underlying C context. */
func (obj *MessengerCloudFsFileDownloadInfo) Ctx() uintptr {
	return uintptr(unsafe.Pointer(obj.cCtx))
}

func NewMessengerCloudFsFileDownloadInfo() *MessengerCloudFsFileDownloadInfo {
	ctx := C.vssq_messenger_cloud_fs_file_download_info_new()
	obj := &MessengerCloudFsFileDownloadInfo{
		cCtx: ctx,
	}
	runtime.SetFinalizer(obj, (*MessengerCloudFsFileDownloadInfo).Delete)
	return obj
}

/* Acquire C context.
* Note. This method is used in generated code only, and SHOULD NOT be used in another way.
 */
func NewMessengerCloudFsFileDownloadInfoWithCtx(pointer unsafe.Pointer) *MessengerCloudFsFileDownloadInfo {
	ctx := (*C.vssq_messenger_cloud_fs_file_download_info_t /*ct2*/)(pointer)
	obj := &MessengerCloudFsFileDownloadInfo{
		cCtx: ctx,
	}
	runtime.SetFinalizer(obj, (*MessengerCloudFsFileDownloadInfo).Delete)
	return obj
}

/* Acquire retained C context.
* Note. This method is used in generated code only, and SHOULD NOT be used in another way.
 */
func NewMessengerCloudFsFileDownloadInfoCopy(pointer unsafe.Pointer) *MessengerCloudFsFileDownloadInfo {
	ctx := (*C.vssq_messenger_cloud_fs_file_download_info_t /*ct2*/)(pointer)
	obj := &MessengerCloudFsFileDownloadInfo{
		cCtx: C.vssq_messenger_cloud_fs_file_download_info_shallow_copy(ctx),
	}
	runtime.SetFinalizer(obj, (*MessengerCloudFsFileDownloadInfo).Delete)
	return obj
}

/*
* Release underlying C context.
 */
func (obj *MessengerCloudFsFileDownloadInfo) Delete() {
	if obj == nil {
		return
	}
	runtime.SetFinalizer(obj, nil)
	obj.delete()
}

/*
* Release underlying C context.
 */
func (obj *MessengerCloudFsFileDownloadInfo) delete() {
	C.vssq_messenger_cloud_fs_file_download_info_delete(obj.cCtx)
}

/*
* Create fully defined object.
 */
func NewMessengerCloudFsFileDownloadInfoWith(link string, fileEncryptedKey []byte) *MessengerCloudFsFileDownloadInfo {
	linkChar := C.CString(link)
	defer C.free(unsafe.Pointer(linkChar))
	linkStr := C.vsc_str_from_str(linkChar)
	fileEncryptedKeyData := helperWrapData(fileEncryptedKey)

	proxyResult := /*pr4*/ C.vssq_messenger_cloud_fs_file_download_info_new_with(linkStr, fileEncryptedKeyData)

	runtime.KeepAlive(link)

	obj := &MessengerCloudFsFileDownloadInfo{
		cCtx: proxyResult,
	}
	runtime.SetFinalizer(obj, (*MessengerCloudFsFileDownloadInfo).Delete)
	return obj
}

func (obj *MessengerCloudFsFileDownloadInfo) Link() string {
	proxyResult := /*pr4*/ C.vssq_messenger_cloud_fs_file_download_info_link(obj.cCtx)

	runtime.KeepAlive(obj)

	return C.GoString(C.vsc_str_chars(proxyResult)) /* r5.1 */
}

func (obj *MessengerCloudFsFileDownloadInfo) FileEncryptedKey() []byte {
	proxyResult := /*pr4*/ C.vssq_messenger_cloud_fs_file_download_info_file_encrypted_key(obj.cCtx)

	runtime.KeepAlive(obj)

	return helperExtractData(proxyResult) /* r1 */
}
