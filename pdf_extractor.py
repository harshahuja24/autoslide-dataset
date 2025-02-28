import pdfplumber
import argparse
import os
import re

def extract_raw_text_from_pdf(pdf_path, output_path):
    """
    Extract raw text from a PDF file and write it to an output file,
    removing extra spaces and newlines.
    
    Args:
        pdf_path (str): Path to the PDF file
        output_path (str): Path to the output text file
    """
    try:
        if not os.path.isfile(pdf_path):
            print(f"Error: PDF file '{pdf_path}' not found.")
            return False
        

        with pdfplumber.open(pdf_path) as pdf:
            all_text = ""
            for page in pdf.pages:
                all_text += page.extract_text()
          
            clean_text = re.sub(r'\s+', " ", all_text)
            
      
            with open(output_path, 'w', encoding='utf-8') as output_file:
                output_file.write(clean_text)
            
            print(f"Raw text extraction complete. Output written to '{output_path}'")
            return True
    
    except Exception as e:
        print(f"An error occurred: {str(e)}")
        return False

def main():
  
    parser = argparse.ArgumentParser(description='Extract raw text from a PDF file')
    parser.add_argument('pdf_path', help='Path to the PDF file')
    parser.add_argument('-o', '--output', help='Path to the output text file (default: output.txt)')
    

    args = parser.parse_args()
    
   
    output_path = args.output if args.output else 'output.txt'
    

    extract_raw_text_from_pdf(args.pdf_path, output_path)

if __name__ == "__main__":
    main()