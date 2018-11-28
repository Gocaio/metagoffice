/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package main

import (
	"fmt"

	"github.com/kevinborras/metagoffice"
)

func main() {

	content := metagoffice.GetContent("example/document.docx")

	fmt.Println("Title: ", content.Title)
	fmt.Println("Subject: ", content.Subject)
	fmt.Println("Creator: ", content.Creator)
	fmt.Println("Keywords: ", content.Keywords)
	fmt.Println("Description: ", content.Description)
	fmt.Println("Last modified by: ", content.LastModifiedBy)
	fmt.Println("Revision: ", content.Revision)
	fmt.Println("Created: ", content.Created)
	fmt.Println("Modified: ", content.Modified)
	fmt.Println("Category: ", content.Category)

}
