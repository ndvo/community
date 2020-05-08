--https://stackoverflow.com/questions/12965463/humanized-or-natural-number-sorting-of-mixed-word-and-number-strings#answer-20667107
CREATE OR REPLACE FUNCTION human_sort(text)
  RETURNS text[] AS
$BODY$   
  /* Split the input text into contiguous chunks where no numbers appear,
     and contiguous chunks of only numbers. For the numbers, add leading 
     zeros to 20 digits, so we can use one text array, but sort the 
     numbers as if they were big integers.

       For example, human_sort('Run 12 Miles') gives
            {'Run ', '00000000000000000012', ' Miles'}
  */
  select array_agg(
    case
      when a.match_array[1]::text is not null 
        then a.match_array[1]::text         
      else lpad(a.match_array[2]::text, 20::int, '0'::text)::text                                      
    end::text)
    from (
      select regexp_matches(
        case when $1 = '' then null else $1 end, E'(\\D+)|(\\d+)', 'g'
      ) AS match_array      
    ) AS a  
$BODY$
  LANGUAGE sql IMMUTABLE;
