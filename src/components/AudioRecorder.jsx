import { useState, useRef, useEffect } from 'react';
import { Button } from './ui/button';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";

export function AudioRecorder() {
  const [devices, setDevices] = useState([]);
  const [selectedDevice, setSelectedDevice] = useState('');
  const [isRecording, setIsRecording] = useState(false);
  const [audioBlob, setAudioBlob] = useState(null);
  const [notes, setNotes] = useState([]);
  const [error, setError] = useState('');

  const mediaRecorderRef = useRef(null);
  const chunksRef = useRef([]);

  useEffect(() => {
    async function getDevices() {
      try {
        const devices = await navigator.mediaDevices.enumerateDevices();
        const audioDevices = devices.filter(device => device.kind === 'audioinput');
        setDevices(audioDevices);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to get audio devices');
      }
    }

    getDevices();
  }, []);

  async function startRecording() {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({
        audio: { deviceId: selectedDevice }
      });

      const mediaRecorder = new MediaRecorder(stream);
      mediaRecorderRef.current = mediaRecorder;
      chunksRef.current = [];

      mediaRecorder.ondataavailable = (event) => {
        chunksRef.current.push(event.data);
      };

      mediaRecorder.onstop = () => {
        const blob = new Blob(chunksRef.current, { type: 'audio/webm' });
        setAudioBlob(blob);
        stream.getTracks().forEach(track => track.stop());
      };

      mediaRecorder.start();
      setIsRecording(true);
      setError('');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to start recording');
    }
  }

  async function stopRecording() {
    try {
      mediaRecorderRef.current?.stop();
      setIsRecording(false);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to stop recording');
    }
  }

  return (
    <div className="space-y-4">
      <Select
        value={selectedDevice}
        onValueChange={setSelectedDevice}
        disabled={isRecording}
      >
        <SelectTrigger>
          <SelectValue placeholder="Select Microphone" />
        </SelectTrigger>
        <SelectContent>
          {devices.map(device => (
            <SelectItem key={device.deviceId} value={device.deviceId}>
              {device.label || `Microphone ${device.deviceId}`}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>

      {error && (
        <div className="text-destructive text-sm">{error}</div>
      )}

      <div className="flex gap-4 justify-center">
        <Button 
          onClick={isRecording ? stopRecording : startRecording}
          variant={isRecording ? "destructive" : "default"}
          disabled={!selectedDevice}
        >
          {isRecording ? 'Stop Recording' : 'Start Recording'}
        </Button>
      </div>

      {audioBlob && (
        <audio 
          controls 
          src={URL.createObjectURL(audioBlob)} 
          className="w-full" 
        />
      )}

      <div className="space-y-2">
        {notes.map((note, index) => (
          <div key={index} className="bg-card p-4 rounded-lg shadow-sm">
            <p className="text-sm">{note.summary}</p>
          </div>
        ))}
      </div>
    </div>
  );
} 