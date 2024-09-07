# LekScrape

## What is it?

Google Docs scraping utility designed to grab and organize the LekMod Master List of Civilizations Google Doc for easy storage in PocketBase. Mainly used in order to assist with developing [LekBot](https://github.com/jacksondarman/lekbot).

## How do I run it?

_NOTE_: In order to run this project, you need a Google Docs API key and a Google Doc that you have permission to read. I can't help you with the latter, but the [former has plenty of resources to get you started](https://developers.google.com/docs/api/how-tos/overview).

- Open a terminal and navigate to the cloned project directory. Create a `.env` file by running `touch .env` or by using another method.
- Create a `GOOGLE_CREDENTIALS_FILE_PATH` field and a `DOCUMENT_ID` field. Add your JSON credentials filepath as a value for `GOOGLE_CREDENTIALS_FILE_PATH`, and the ID of the Google Doc you want to scrape in `DOCUMENT_ID`.
- Run the project!
