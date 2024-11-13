import { useState, useRef, useEffect } from 'react'
import { Button } from './components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from './components/ui/card'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./components/ui/select"
import { AudioRecorder } from './components/AudioRecorder'

export default function App() {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center p-4">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle>Audio Notes</CardTitle>
        </CardHeader>
        <CardContent>
          <AudioRecorder />
        </CardContent>
      </Card>
    </div>
  )
}