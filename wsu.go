package wsu

import (
	"encoding/xml"

	"github.com/gosoap/xsd"
)

//
// This type defines the fault code value for Timestamp message expiration.
//
// type TTimestampFault *QName

// const (
// 	TTimestampFaultWsuMessageExpired TTimestampFault = "wsu:MessageExpired"
// )

const Namespace = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"

type CommonAttrs struct {
	ID *xsd.ID `xml:"xmlns:wsu,attr,omitempty"`
}

// Timestamp element allows Timestamps to be applied anywhere element wildcards are present,
// including as a SOAP header.
type Timestamp TimestampType

// Expires element allows an expiration time to be applied anywhere element wildcards are present.
type Expires struct {
	XMLName xml.Name `xml:"wsu:Expires"`
	NS      string   `xml:"xmlns:wsu,attr"`

	ID *xsd.ID `xml:",attr,omitempty"`

	Value string `xml:",chardata"`
}

// Created element allows a creation time to be applied anywhere element wildcards are present.
//
// <xsd:element name="Created" type="wsu:AttributedDateTime">
//  <xsd:annotation>
//   <xsd:documentation>
// This element allows a creation time to be applied anywhere element wildcards are present.
// 	 </xsd:documentation>
//  </xsd:annotation>
// </xsd:element>
type Created struct {
	XMLName xml.Name `xml:"wsu:Created"`
	NS      string   `xml:"xmlns:wsu,attr"`

	ID *xsd.ID `xml:",attr,omitempty"`

	Value string `xml:",chardata"`
}

// AttributedURI type is for elements whose [children] is an anyURI and can have arbitrary attributes.
//
// <xsd:complexType name="AttributedURI">
// 	<xsd:annotation>
// 	 <xsd:documentation>
// This type is for elements whose [children] is an anyURI and can have arbitrary attributes.
//   </xsd:documentation>
// 	</xsd:annotation>
// 	<xsd:simpleContent>
// 	 <xsd:extension base="xsd:anyURI">
// 	  <xsd:attributeGroup ref="wsu:commonAtts"/>
// 	 </xsd:extension>
// 	</xsd:simpleContent>
// </xsd:complexType>
type AttributedURI struct {
	ID *xsd.ID `xml:"xmlns:wsu,attr,omitempty"`

	Value *xsd.AnyURI `xml:",chardata"`
}

// TimestampType complex type ties together the timestamp related elements into a composite type.
type TimestampType struct {
	XMLName xml.Name `xml:"wsu:Timestamp"`

	ID *xsd.ID `xml:"xmlns:wsu,attr,omitempty"`

	Created *Created `xml:"Created,omitempty"`

	Expires *Expires `xml:"Expires,omitempty"`
}
