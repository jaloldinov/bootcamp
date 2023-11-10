create or replace function multiply(x int,y int)
   returns int 
   language plpgsql
  as
$$
begin
 return x*y;
end;
$$;



create or replace function print1to9()
   returns int[] 
   language plpgsql
  as
$$
declare 
i int=0;
nums int[];
begin
 loop 
    i=i+1;
    nums:=ARRAY_APPEND(nums,i);
    if i>10 then
    exit;
    end if;
 end loop;

 return nums;
end;
$$;