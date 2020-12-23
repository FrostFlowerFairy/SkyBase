<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * AuthTarget is the payload/input to Authorizer expressions
 *
 * Generated from protobuf message <code>api.AuthTarget</code>
 */
class AuthTarget extends \Google\Protobuf\Internal\Message
{
    /**
     * user is the user making the request
     *
     * Generated from protobuf field <code>.api.Doc user = 1 [(.validator.field) = {</code>
     */
    private $user = null;
    /**
     * target is the request/response represented as a Struct
     *
     * Generated from protobuf field <code>.google.protobuf.Struct target = 2;</code>
     */
    private $target = null;
    /**
     * headers are the request headers
     *
     * Generated from protobuf field <code>map<string, string> headers = 3;</code>
     */
    private $headers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Doc $user
     *           user is the user making the request
     *     @type \Google\Protobuf\Struct $target
     *           target is the request/response represented as a Struct
     *     @type array|\Google\Protobuf\Internal\MapField $headers
     *           headers are the request headers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * user is the user making the request
     *
     * Generated from protobuf field <code>.api.Doc user = 1 [(.validator.field) = {</code>
     * @return \Api\Doc
     */
    public function getUser()
    {
        return $this->user;
    }

    /**
     * user is the user making the request
     *
     * Generated from protobuf field <code>.api.Doc user = 1 [(.validator.field) = {</code>
     * @param \Api\Doc $var
     * @return $this
     */
    public function setUser($var)
    {
        GPBUtil::checkMessage($var, \Api\Doc::class);
        $this->user = $var;

        return $this;
    }

    /**
     * target is the request/response represented as a Struct
     *
     * Generated from protobuf field <code>.google.protobuf.Struct target = 2;</code>
     * @return \Google\Protobuf\Struct
     */
    public function getTarget()
    {
        return $this->target;
    }

    /**
     * target is the request/response represented as a Struct
     *
     * Generated from protobuf field <code>.google.protobuf.Struct target = 2;</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setTarget($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->target = $var;

        return $this;
    }

    /**
     * headers are the request headers
     *
     * Generated from protobuf field <code>map<string, string> headers = 3;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getHeaders()
    {
        return $this->headers;
    }

    /**
     * headers are the request headers
     *
     * Generated from protobuf field <code>map<string, string> headers = 3;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setHeaders($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->headers = $arr;

        return $this;
    }

}

