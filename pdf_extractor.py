import pdfplumber
import argparse
import sys
import re
import os

def extract_raw_text_from_pdf(pdf_path, output_path=None):
    """
    Extract raw text from a PDF file and write it to an output file,
    removing extra spaces and newlines.
    
    Args:
        pdf_path (str): Path to the PDF file
        output_path (str, optional): Path to the output text file
    """
    try:
        if not os.path.isfile(pdf_path):
            sys.stderr.write(f"Error: PDF file '{pdf_path}' not found.\n")
            return False
        
        with pdfplumber.open(pdf_path) as pdf:
            all_text = ""
            for page in pdf.pages:
                text = page.extract_text()
                if text:
                    all_text += text + " "
          
            # Clean the text by removing excessive whitespace
            clean_text = re.sub(r'\s+', " ", all_text).strip()
            
            # If output_path is provided, write to file
            if output_path:
                with open(output_path, 'w', encoding='utf-8') as output_file:
                    output_file.write(clean_text)
                sys.stderr.write(f"Raw text extraction complete. Output written to '{output_path}'\n")
            
            # For Go integration: write to temp file and return the path
            temp_file = "temp_pdf_text.txt"
            with open(temp_file, 'w', encoding='utf-8') as f:
                f.write(clean_text)
            
            return True
    
    except Exception as e:
        sys.stderr.write(f"PDF extraction error: {str(e)}\n")
        return False

def main():
    parser = argparse.ArgumentParser(description='Extract raw text from a PDF file')
    parser.add_argument('pdf_path', help='Path to the PDF file')
    parser.add_argument('-o', '--output', help='Path to the output text file (optional)')
    
    args = parser.parse_args()
    
    # Extract and write to file
    success = extract_raw_text_from_pdf(args.pdf_path, args.output)
    
    # Exit with appropriate status code
    sys.exit(0 if success else 1)

if __name__ == "__main__":
    main()