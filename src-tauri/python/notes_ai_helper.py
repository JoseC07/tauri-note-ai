import sys
import json
import os
from pathlib import Path
from typing import Dict, Any

import torch
from transformers import pipeline, AutoTokenizer, AutoModelForSequenceClassification
import sounddevice as sd
import numpy as np
import librosa
from dotenv import load_dotenv

class NotesAIHelper:
    def __init__(self):
        # Load environment variables
        load_dotenv()
        
        # Initialize models directory
        self.models_dir = Path(os.getenv('APP_DATA_DIR')) / 'models'
        self.models_dir.mkdir(parents=True, exist_ok=True)
        
        # Initialize AI models
        self.summarizer = pipeline('summarization', 
                                 model='facebook/bart-large-cnn',
                                 device=0 if torch.cuda.is_available() else -1)
        
        self.classifier = pipeline('text-classification',
                                 model='distilbert-base-uncased-finetuned-sst-2-english',
                                 device=0 if torch.cuda.is_available() else -1)
        
        # Audio processing settings
        self.sample_rate = 16000
    
    def summarize(self, text: str) -> str:
        """Generate a summary of the input text."""
        try:
            summary = self.summarizer(text, max_length=130, min_length=30, do_sample=False)
            return summary[0]['summary_text']
        except Exception as e:
            return f"Error generating summary: {str(e)}"
    
    def classify(self, text: str) -> Dict[str, Any]:
        """Classify the sentiment/topic of the text."""
        try:
            result = self.classifier(text)
            return result[0]
        except Exception as e:
            return {"error": f"Classification failed: {str(e)}"}
    
    def process_audio(self, audio_path: str) -> str:
        """Convert audio to text."""
        try:
            # Load audio file
            audio, _ = librosa.load(audio_path, sr=self.sample_rate)
            
            # TODO: Implement speech-to-text conversion
            # This is a placeholder - you'll need to implement actual STT
            return "Audio processing not yet implemented"
            
        except Exception as e:
            return f"Error processing audio: {str(e)}"
    
    def process_command(self, command: str, input_data: str) -> Dict[str, Any]:
        """Process commands from the Go backend."""
        if command == 'summarize':
            return {"result": self.summarize(input_data)}
        elif command == 'classify':
            return self.classify(input_data)
        elif command == 'process_audio':
            return {"result": self.process_audio(input_data)}
        else:
            return {"error": "Unknown command"}

def main():
    """Main entry point for the script."""
    if len(sys.argv) < 3:
        print(json.dumps({"error": "Insufficient arguments"}))
        sys.exit(1)
    
    try:
        helper = NotesAIHelper()
        command = sys.argv[1]
        input_data = sys.argv[2]
        result = helper.process_command(command, input_data)
        print(json.dumps(result))
    except Exception as e:
        print(json.dumps({"error": f"Processing failed: {str(e)}"}))
        sys.exit(1)

if __name__ == "__main__":
    main()
