


let data =[5, 9, 13, 17,45,67,89,100];


let find=89
let start =0;
let end=data.length-1;
let position


while(end>=start)
    {
        let mid=Math.floor((start+end)/2)

        if(data[mid]==find)
            {
                position=mid
                break;
            }
            else if(data[mid] >find)
                {
                    end=mid-1
                }
                else{
                    start=mid+1
                }
    }


    position = position || `find value not found`


    console.log(position);