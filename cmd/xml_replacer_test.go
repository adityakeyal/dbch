package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/beevik/etree"
)

func TestXMLReplacer(t *testing.T) {

	x := `<?xml version="1.0" encoding="UTF-8"?>

		<!-- $Id: xenos-jms-context.xml,v 1.6.14.1 2017/05/09 10:29:15 shivams Exp $ -->
		<!-- $Revision: 1.6.14.1 $ -->
		<!-- $Date: 2017/05/09 10:29:15 $ -->
		
		
		<!--
			Application context definition for JMS in non-XA setting.
		-->
		<beans xmlns="http://www.springframework.org/schema/beans"
						xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
						xmlns:context="http://www.springframework.org/schema/context"
						xmlns:p="http://www.springframework.org/schema/p"
						xsi:schemaLocation="
								http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-3.0.xsd
								http://www.springframework.org/schema/context http://www.springframework.org/schema/context/spring-context-3.0.xsd">
		
		
			<!-- ========================= RESOURCE DEFINITIONS ========================= -->
		
			<!-- <context:property-placeholder location="classpath:META-INF/spring/xenos-jms.properties"/>-->
			<bean 	id="jmsproperties" 
					class="org.springframework.beans.factory.config.PropertyPlaceholderConfigurer">
				<property name="locations">
					<list>
						<value>classpath:META-INF/spring/xenos-jms.properties</value>
					</list>
				</property>
				<property name="ignoreUnresolvablePlaceholders" value="true"/>
			</bean>
		
			<bean 	id="jms-config"
					class="com.nri.xenos.inf.JmsConfigImpl"
					init-method="init" />
		
			<bean	id="jms-connection-factory"
					class="com.nri.xenos.inf.JmsConnectionFactoryImpl">
				<constructor-arg ref="xenosQueueConnectionFactoryBean" />
			</bean>
		
		
			<bean	id="com.ibm.msg.client.wmq.common.CommonConstants.WMQ_CM_CLIENT"
					class="org.springframework.beans.factory.config.FieldRetrievingFactoryBean" />
		
			<bean	id="org.apache.commons.pool.impl.GenericObjectPool.WHEN_EXHAUSTED_GROW"
					class="org.springframework.beans.factory.config.FieldRetrievingFactoryBean" />
		
			<bean	id="xenosQueueConnectionFactoryBean"
					class="com.nri.xenos.inf.XenosQueueConnectionFactoryBean"
					init-method="init"
					p:userId = "${jms.queueManager.userid}"
					p:password = "${jms.queueManager.password}"
					destroy-method="destroy">
				<constructor-arg ref="mqQueueConnectionFactory" />
				<constructor-arg ref="jms-config" />
				<constructor-arg ref="org.apache.commons.pool.impl.GenericObjectPool.WHEN_EXHAUSTED_GROW" />
			</bean>
		
			<!-- bean	id="mqQueueConnectionFactory"
					class="com.ibm.mq.jms.MQQueueConnectionFactory"
					p:queueManager="${jms.queueManager.name}"
					p:transportType-ref="com.ibm.msg.client.wmq.common.CommonConstants.WMQ_CM_CLIENT"
					p:hostName="${jms.queueManager.host}"
					p:port="${jms.queueManager.port}"
					p:channel="${jms.queueManager.channel}" /-->
			
			<bean  id="mqQueueConnectionFactory"
										  class="org.apache.activemq.ActiveMQConnectionFactory">
				 <property name = "brokerURL">
					<value>tcp://localhost:61616?wireFormat.maxInactivityDuration=1000</value>
				 </property>
				 </bean>
				 
				 <!-- For Active MQ Failover -->
				 <!-- bean  id="mqQueueConnectionFactory"
										  class="org.apache.activemq.ActiveMQConnectionFactory">
				 <property name = "brokerURL">
					  <value>failover:(tcp://172.16.30.187:61616?wireFormat.maxInactivityDuration=1000,tcp://172.16.30.188:61616?wireFormat.maxInactivityDuration=1000)</value>
					  
					  http://:8161/
					  
				 </property>
				 </bean -->
		
		</beans>
		`

	//source := "//beans/bean/property/value"
	// targetValue := "tcp://adityak:61616?daemon=true&amp;jms.redeliveryPolicy.maximumRedeliveries=-1"

	doc := etree.NewDocument()
	err := doc.ReadFromString(x)
	if err != nil {
		panic(err)
	}

	//beans/bean[id=mqQueueConnectionFactory]/property[name=brokerURL]/value

	for _, t := range doc.FindElements("//bean[@id='mqQueueConnectionFactory']/property[@name='brokerURL']/value") {
		t.SetText("Aditya")
	}

	str, _ := doc.WriteToString()

	fmt.Println(str)

	//identifyTag(doc.Root(), source, "Aditya")

}

func identifyTag(element *etree.Element, xpath string, value string) {

	tagName := updateAndReturnXpath(&xpath)
	if element.Tag == tagName {
		if len(element.ChildElements()) == 0 {
			fmt.Println(*element)
		} else {
			for _, val := range element.ChildElements() {
				identifyTag(val, xpath, value)
			}
		}
	}
}

func updateAndReturnXpath(xpath *string) string {

	tags := strings.Split(*xpath, "/")

	for i := 0; i < len(tags); i++ {
		if tags[i] != "" {
			*xpath = strings.Join(tags[(i+1):], "/")
			return tags[i]
		}
	}
	return ""
}
