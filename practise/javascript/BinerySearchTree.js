




let data={
    value:null,
    left:null,
    right:null
}




    function addData(val){
        if(data.value==null)
            {
                data.value=val
            }
            else{
                insertNode(data,val)
            }

    }

    function insertNode(root,val){
       
            if(root.value>val)
                {
                    if(root.left==null)
                        {
                            let data={
                                value:val,
                                left:null,
                                right:null
                            }
                            root.left=data
                        }
                        else{
                            insertNode(root.left,val)
                        }
                }
                else{
                    if(root.right==null)
                        {
                            let data={
                                value:val,
                                left:null,
                                right:null
                            }
                            root.right=data

                            // root.value=val
                            // root.left=null
                            // root.right=null
                        }else{
                            insertNode(root.right,val)
                        }

                }

    }


        addData(100)
        addData(90)
        addData(80)
        addData(85)
        addData(79)
        addData(110)
        addData(91)






        console.log(JSON.stringify(data));
