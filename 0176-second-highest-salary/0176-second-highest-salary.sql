# Write your MySQL query statement below
select(
    select distinct salary  
    from Employee 
    # where salary not in (
    #    select salary from Employee
    #    order by salary desc
    #    limit 1
    # ) # This version of MySQL doesn't yet support 'LIMIT & IN/ALL/ANY/SOME subquery'
    order by salary desc
    limit 1 
    offset 1
) as SecondHighestSalary