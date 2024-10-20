




function checkAnagram(str1,str2)
{
    if(str1.length!=str2.length)
        {
            return false
        }

       let da={};

       for( ab of str1)
        {
            da[ab]= (da[ab] || 0) + 1
        }
        console.log(da)

        for(let ab of str2)
            {
                if (da[ab]){
                    da[ab]=da[ab]-1
                }   
                else 
                {
                    da[ab]= (da[ab] || 0) + 1
                }
            }

            for(let ab in da)
                {
                    if(da[ab]>0)
                        {
                            return false
                        }
                }
       
console.log(da)
                return true;
}




console.log(checkAnagram("listen","silent"))