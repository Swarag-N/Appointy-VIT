# Initial Planing

Required Models

- Users (Partipants/ Particpating)
  - Name            String
  - Email           String
  - Passphrase      String (optional)
  - Future Meetings []Meetings

- Meetings
  - Title           String
  - Start Time      Timestamp Integer
  - End Time        Timestamp Integer
  - Created Time    Timestamp Integer
  - Participants    []Users
    - User ID
    - email
    - RSVP
      - Yes
      - No
      - Maybe Yes
      - No Answered


  - Meeting ID
  - User ID
  - RSVP
    - Yes
    - No
    - Maybe Yes
    - No Answered

Add

- Lock Condition
- MongoDB

Routes

- Querying
  - Pagination
    - Offset
    - limit
  - Filter
    - start
    - end

- RSVP Confirmation or Meeting Create Validator
  - GET User
    - Populate Meetings with []FutureMeeting (where RSVP == "YES")
    - Set StartTime and EndTime and make an Array[2] `[startTime, endTIme]`
    - Form meetingsArray[n] `[[startTime, endTIme],[startTime, endTIme]]`
  - from Post Data
    - Convert Dates to Timestamps
  - From User data and Post Data
    - if any of `StartTime` or `Endtime` Lies between `[startTime, endTIme]`
      - reject



