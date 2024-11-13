use tauri::{command, Runtime};
use serde::Serialize;
use tauri::Builder as TauriBuilder;
use std::thread::Builder as ThreadBuilder;
#[derive(Debug, Serialize)]
struct AudioPermissionResponse {
    granted: bool,
    error: Option<String>,
}

#[command]
fn request_microphone_permission<R: Runtime>(
    _window: tauri::Window<R>
) -> Result<AudioPermissionResponse, String> {
    #[cfg(target_os = "macos")]
    {
        use coreaudio::sys::{
            AudioObjectGetPropertyData,
            AudioObjectPropertyAddress,
            kAudioHardwarePropertyDevices,
            kAudioObjectSystemObject,
            kAudioObjectPropertyScopeGlobal,
            kAudioObjectPropertyElementMaster,
        };
        use std::ptr::null_mut;

        let property_address = AudioObjectPropertyAddress {
            mSelector: kAudioHardwarePropertyDevices,
            mScope: kAudioObjectPropertyScopeGlobal,
            mElement: kAudioObjectPropertyElementMaster,
        };

        let mut data_size = 0u32;
        let status = unsafe {
            AudioObjectGetPropertyData(
                kAudioObjectSystemObject,
                &property_address,
                0,
                null_mut(),
                &mut data_size,
                null_mut(),
            )
        };

        if status == 0 {
            Ok(AudioPermissionResponse { granted: true, error: None })
        } else {
            Ok(AudioPermissionResponse { 
                granted: false, 
                error: Some("Microphone access denied".to_string()) 
            })
        }
    }

    #[cfg(target_os = "windows")]
    {
        Ok(AudioPermissionResponse { granted: true, error: None })
    }

    #[cfg(target_os = "linux")]
    {
        Ok(AudioPermissionResponse { granted: true, error: None })
    }
}

fn main() {
    TauriBuilder::default()
        .invoke_handler(generate_handler![request_microphone_permission])
        .run(generate_context!())
        .expect("error while running Tauri application");
}
