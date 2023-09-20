-- BLOCK --
--info,warning,notice
DO $$
DECLARE 
    a int=5;
    b int=6;
BEGIN
    raise notice '% ', a*b;
END $$;


--  IF --
DO $$
DECLARE 
    x integer=5;
BEGIN
IF x=0 THEN
    raise notice 'x is zero';
ELSIF x%2=1 THEN
    raise notice '% - toq',x;
ELSE 
    raise notice '% - juft',x;
END IF;
END $$;

-- CASE --
DO $$
DECLARE 
    x integer=5;
BEGIN
CASE x%2   
   WHEN 0 THEN raise notice '% - juft',x;  
   WHEN 1 THEN raise notice '% - toq',x;
END CASE;
END $$;

-- LOOP --

DO $$
DECLARE
   i int =0;
BEGIN
    LOOP 
        i=i+1;
        raise info 'i= %',i;
        if i=10 then exit;
        end if;
    END LOOP;
END $$;

-- FOR --

do $$
declare
a int =0;
b int =20;
begin
   for counter in a..b by 5 
   loop
	raise notice 'counter: %', counter;
   end loop;
end; $$;


do $$
begin
   for counter in reverse 10..1 by 2 loop
	raise notice 'counter: %', counter;
   end loop;
end $$;
-- FOR over select result --
do $$
DECLARE
branch RECORD;  -- branches:=[]branch{}
begin
   for branch in SELECT id,name 
                FROM branches
    loop
	raise notice 'branch: id= %  name= %', branch.id,branch.name;
   end loop;
   for _,branch:=range branches{
    fmt.Printf('branch: id= %d  name= %s\n', branch.id,branch.name)
   }
end $$;